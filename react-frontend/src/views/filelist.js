import {Component} from "react";
import {Button, Card, Dropdown, ListGroup, Pagination, Table} from "flowbite-react";
import {ListGroupItem} from "flowbite-react/lib/esm/components/ListGroup/ListGroupItem";
import {Auth0AddQue} from "../auth0";
import {apiClient} from "../api";
import {FileInfo as APIFileInfo} from "ts-cloudwpss23-openapi-cyan";
import {FileInfo} from "ts-cloudwpss23-openapi-cyan";
import * as $ from 'jquery';
import {TableHead} from "flowbite-react/lib/esm/components/Table/TableHead";
import {TableHeadCell} from "flowbite-react/lib/esm/components/Table/TableHeadCell";
import {TableBody} from "flowbite-react/lib/esm/components/Table/TableBody";
import {TableRow} from "flowbite-react/lib/esm/components/Table/TableRow";
import {TableCell} from "flowbite-react/lib/esm/components/Table/TableCell";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faDownload, faTrash, faPlus, faSpinner, faFile, faFolder, faLanguage} from "@fortawesome/free-solid-svg-icons";
import {withAuth0} from "@auth0/auth0-react";


export const Filelist = withAuth0(class extends Component {

    constructor(props) {
        super(props);
        this.onDropdownClick = this.onDropdownClick.bind(this)
        this.onUploadClick = this.onUploadClick.bind(this)
        this.updateFileList = this.updateFileList.bind(this)
        this.deleteFile = this.deleteFile.bind(this)
        this.downloadFile = this.downloadFile.bind(this)
        this.setFileDownloadState = this.setFileDownloadState.bind(this)
        this.translateFile = this.translateFile.bind(this)
        this.setFileTranslateState = this.setFileTranslateState.bind(this)
        let {Bucket} = props;
        this.state = {
            Bucket: Bucket,
            Files: [],
            FileState: {},
            Folders: [],
            ItemsPerPage: 10,
            Uploading: false,
            CurrentPage: 1
        }
        this.updateFileList()
    }

    updateFileList(){
        let parent = this;
        Auth0AddQue(() => {
            $.when(apiClient.getV1Files(parent.state.Bucket)).then(function (status) {
                let response: Array<APIFileInfo> = status.body;
                if (response != null) {
                    let folders = response.filter((f) => f.size === undefined)
                    folders = folders.map((f) => {
                        f.size = 0;
                        return f;
                    })
                    let files = response.filter((f) => f.size > 0)
                    let state = {}
                    files.forEach((f: APIFileInfo) => {
                        state[f.file.name] = {Translate: false, Download: false}
                    })
                    parent.setState({Files: files, Folders: folders, FileState: state})
                } else {
                    parent.setState({Files: [], Folders: [], FileState: {}})
                }

            })
        })
    }

    async onUploadClick(e){
        const pickerOpts = {
            excludeAcceptAllOption: false,
            multiple: false,
        };

        let list = this
        window.showOpenFilePicker(pickerOpts).then((fileHandles: FileSystemFileHandle[]) => {
            list.setState({Uploading: true})
            fileHandles.forEach((fileHandle) => {
                fileHandle.getFile().then((file) => {
                    $.when(apiClient.putV1FileName(list.state.Bucket, file.name, file)).then((response) => {
                        list.setState({Uploading: false})
                        list.updateFileList()
                    }, (reject) => {
                        console.log("Rejected", reject);
                        list.setState({Uploading: false})
                    })
                })
            })
        })
    }
    deleteFile(filename){
        let list = this
        $.when(apiClient.deleteV1FileName(list.state.Bucket, filename)).then((state) => {
            list.updateFileList()
        })
    }

    setFileDownloadState(filename :string, downloading:boolean){
        let fstate = this.state.FileState
        fstate[filename].Download = downloading
        this.setState({FileState: fstate})
    }

    setFileTranslateState(filename :string, translating:boolean){
        let fstate = this.state.FileState
        fstate[filename].Translate = translating
        this.setState({FileState: fstate})
    }

    downloadFile(filename){
        let list = this;
        list.setFileDownloadState(filename, true)
        $.when(apiClient.getV1FileName(list.state.Bucket, filename)).then((state) => {
            let blob :File = state.body;
            // Create blob link to download
            const url = window.URL.createObjectURL(
                new Blob([blob]),
            );
            const link = document.createElement('a');
            link.href = url;
            link.setAttribute(
                'download',
                filename,
            );

            // Append to html link element page
            document.body.appendChild(link);

            // Start download
            link.click();

            // Clean up and remove the link
            link.parentNode.removeChild(link);
            list.setFileDownloadState(filename, false)
        })
    }

    onDropdownClick(e){
        this.setState({ItemsPerPage: e})
        console.log(e);
    }

    translateFile(filename){
        this.setFileTranslateState(filename, true)
        let list = this
        $.when(apiClient.getV1BucketBucketNameTranslateFileName(this.props.Bucket, filename)).then((s) => {
            console.log(s)
            list.setFileTranslateState(filename, false)
            this.updateFileList()
        })
    }
    render() {
        let tableList = this.state.Folders.concat(this.state.Files)
        console.log(tableList)
        let isLoading = this.state.Uploading
        return (
            <div className="pl-5 pr-5 pb-5 w-full">
                <Card className="shadow sticky h-22 mt-5 top-8 mb-5 z-20">
                    <div className="flex items-center justify-center text-center">
                        <div className="w-64" >
                            {this.state.Bucket !== "" && (
                                <h6 className="text-xl tracking-tight text-gray-900 dark:text-white">
                                    <b>Bucket:</b> <i>{this.state.Bucket}</i>
                                </h6>
                            )}
                            {this.state.Bucket === "" &&(
                                <h6 className="text-xl font-bold tracking-tight text-gray-900 dark:text-white">
                                    Bucket: <div className="h-2 bg-slate-500 rounded" />
                                </h6>
                            )}
                        </div>
                        <Pagination
                            currentPage={this.state.CurrentPage}
                            layout="table"
                            showIcons={true}
                            onPageChange={(page) => {this.setState({CurrentPage: page})}}
                            totalPages={Math.round(tableList.length / this.state.ItemsPerPage)}
                            className="pl-5 pr-5"
                        />
                        <div className="grow dark:text-white">
                            <Dropdown label={this.state.ItemsPerPage} size="sm" inline={true} className="dark:text-white z-20">
                                <Dropdown.Item onClick={()=>this.onDropdownClick(10)}>
                                    10
                                </Dropdown.Item>
                                <Dropdown.Item onClick={()=>this.onDropdownClick(25)}>
                                    25
                                </Dropdown.Item>
                                <Dropdown.Item onClick={()=>this.onDropdownClick(50)}>
                                    50
                                </Dropdown.Item>
                                <Dropdown.Item onClick={()=>this.onDropdownClick(100)}>
                                    100
                                </Dropdown.Item>
                            </Dropdown>
                        </div>
                        <div className="pl-5">
                            {!isLoading && (
                                <button onClick={this.onUploadClick} className="aspect-square bg-gradient-to-r from-rose-400 to-pink-700 rounded-full w-10">
                                    <FontAwesomeIcon color="white" icon={faPlus} />
                                </button>
                            )}
                            {isLoading && (
                                <button className="aspect-square bg-gradient-to-r from-rose-400 to-pink-700 rounded-full w-10" disabled={true}>
                                    <FontAwesomeIcon color="white" icon={faSpinner} className="animate-spin" />
                                </button>
                            )}
                        </div>
                    </div>
                </Card>
                <Table className="shadow -z-10" striped={true}>
                    <TableHead className="sticky top-0">
                        <TableHeadCell>Type</TableHeadCell>
                        <TableHeadCell>Name</TableHeadCell>
                        <TableHeadCell>Size (kB)</TableHeadCell>
                        <TableHeadCell>Last Modified</TableHeadCell>
                        <TableHeadCell>Translate</TableHeadCell>
                        <TableHeadCell>Download</TableHeadCell>
                        <TableHeadCell>Delete</TableHeadCell>
                    </TableHead>
                    <TableBody className="bg-neutral-200">
                        {this.state.Files >= 0 && ([0,0,0,0,0,0,0,0,0,0]).map((i) => {
                            return (
                                <TableRow className="animate-pulse">
                                    <TableCell><div className="h-2 bg-slate-500 rounded-full w-2"></div></TableCell>
                                    <TableCell><div className="h-2 bg-slate-500 rounded"></div></TableCell>
                                    <TableCell><div className="h-2 bg-slate-500 rounded"></div></TableCell>
                                    <TableCell><div className="h-2 bg-slate-500 rounded"></div></TableCell>
                                    <TableCell><div className="h-2 bg-slate-500 rounded"></div></TableCell>
                                    <TableCell><div className="h-2 bg-slate-500 rounded"></div></TableCell>
                                    <TableCell><div className="h-2 bg-slate-500 rounded"></div></TableCell>
                                </TableRow>
                            )
                        })}
                    {tableList.length > 0 && tableList.map((file: FileInfo, index) => {
                        if((index >= ((this.state.CurrentPage-1) * this.state.ItemsPerPage)) &&
                            index < this.state.CurrentPage * this.state.ItemsPerPage){
                            let isFolder = file.size === 0;
                            let isPDF = file.file.name.endsWith(".pdf") || file.file.name.endsWith(".docx")
                            let icon = isFolder ? faFolder : faFile
                        return (
                            <TableRow>
                                <TableCell><FontAwesomeIcon icon={icon}/></TableCell>
                                <TableCell>{file.file.name}</TableCell>
                                <TableCell>{Math.round(file.size/1024)}</TableCell>
                                <TableCell>{file.lastmod}</TableCell>
                                <TableCell align={"right"} width={"3rem"}>{isPDF && (<Button onClick={() => {this.translateFile(file.file.name)}} className={`border-0 bg-gradient-to-r from-blue-400 to-cyan-700`}><FontAwesomeIcon icon={this.state.FileState[file.file.name].Translate ? faSpinner : faLanguage} className={this.state.FileState[file.file.name].Translate ? "animate-spin" : ""}/></Button>)}</TableCell>
                                <TableCell align={"right"} width={"3rem"}>{!isFolder && (<Button onClick={() => {this.downloadFile(file.file.name)}} className={`border-0 bg-gradient-to-r from-blue-400 to-cyan-700`}><FontAwesomeIcon icon={this.state.FileState[file.file.name].Download ? faSpinner : faDownload} className={this.state.FileState[file.file.name].Download ? "animate-spin" : ""}/></Button>)}</TableCell>
                                <TableCell align={"right"} width={"3rem"}><Button onClick={() => {this.deleteFile(file.file.name)}} className="border-0 bg-gradient-to-r from-red-400 to-rose-800"><FontAwesomeIcon icon={faTrash} /></Button></TableCell>
                            </TableRow>
                        )}
                    })}
                    </TableBody>
                </Table>
            </div>
        );
    }
})