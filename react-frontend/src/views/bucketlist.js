import {Component} from "react";
import {withAuth0} from "@auth0/auth0-react";
import {Button, Card, Dropdown, Pagination, Table} from "flowbite-react";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faDownload, faFile, faFolder, faPlus, faSpinner, faTrash, faCircleCheck as faCircleCheckSolid} from "@fortawesome/free-solid-svg-icons";
import {faCircleCheck as faCircleCheckRegular} from "@fortawesome/free-regular-svg-icons";
import {TableHead} from "flowbite-react/lib/esm/components/Table/TableHead";
import {TableHeadCell} from "flowbite-react/lib/esm/components/Table/TableHeadCell";
import {TableBody} from "flowbite-react/lib/esm/components/Table/TableBody";
import {TableRow} from "flowbite-react/lib/esm/components/Table/TableRow";
import {TableCell} from "flowbite-react/lib/esm/components/Table/TableCell";
import {FileInfo} from "ts-cloudwpss23-openapi-cyan";
import {BucketInfo} from "ts-cloudwpss23-openapi-cyan";
import {Bucket} from "ts-cloudwpss23-openapi-cyan";
import {Auth0AddQue} from "../auth0";
import * as $ from "jquery";
import {apiClient} from "../api";
import {PostV1BucketNameRequest} from "ts-cloudwpss23-openapi-cyan";

export const BucketList = withAuth0(class extends Component {
    constructor(props) {
        super(props);
        this.onDropdownClick = this.onDropdownClick.bind(this)
        this.onCreateClick = this.onCreateClick.bind(this)
        this.updateBucketList = this.updateBucketList.bind(this)
        this.deleteBucket = this.deleteBucket.bind(this)

        this.state = {
            ItemsPerPage: 10,
            CurrentPage: 1,
            Buckets: [],
            Creating: false
        }

        this.updateBucketList()
    }

    updateBucketList(){
        let list = this
        Auth0AddQue(() => {
            $.when(apiClient.getV1Buckets()).then((state) => {
                console.log(state)
                let buckets :Array<Bucket> = state.body;
                list.setState({Buckets: buckets})
            })
        })
    }

    onDropdownClick(e){
        this.setState({ItemsPerPage: e})
    }

    onCreateClick(e){
        let name = prompt("Bitte den Namen des neuen Buckets angeben", "");
        if (name !== "") {
            name = name.toLowerCase()
            name = name.replace(" ", "-")
            let list = this
            list.setState({Creating: true})
            let req : PostV1BucketNameRequest = {name: name}
            $.when(apiClient.postV1BucketName(name, req)).then((s) => {
                console.log(s)
                list.setState({Creating: false})
                list.updateBucketList()
            })
        }
    }

    deleteBucket(bucketName) {
        let list = this
        $.when(apiClient.deleteV1BucketName(bucketName)).then(() => {
            list.updateBucketList()
        })
    }

    render() {
        const { changeBucketHandler, Bucket } = this.props
        let isLoading = this.state.Creating
        let tableList : Array<Bucket> = this.state.Buckets
        return (
            <div className="pl-5 pr-5 pb-5 w-full">
                <Card className="shadow sticky h-22 mt-5 top-8 mb-5 z-20">
                    <div className="flex items-center justify-center text-center">
                        <div className="w-64">
                            {Bucket !== "" && (
                                <h6 className="text-xl tracking-tight text-gray-900 dark:text-white">
                                    <b>Current Bucket:</b> <i>{Bucket}</i>
                                </h6>
                            )}
                            {Bucket === "" && (
                                <h6 className="text-xl font-bold tracking-tight text-gray-900 dark:text-white">
                                    Current Bucket: <div className="h-2 bg-slate-500 rounded"/>
                                </h6>
                            )}
                        </div>
                        <Pagination
                            currentPage={this.state.CurrentPage}
                            layout="table"
                            showIcons={true}
                            onPageChange={(page) => {
                                this.setState({CurrentPage: page})
                            }}
                            totalPages={tableList != null ? Math.round(tableList.length / this.state.ItemsPerPage) + 1 : 1}
                            className="pl-5 pr-5"
                        />
                        <div className="grow dark:text-white">
                            <Dropdown label={this.state.ItemsPerPage} size="sm" inline={true}
                                      className="dark:text-white z-20">
                                <Dropdown.Item onClick={() => this.onDropdownClick(10)}>
                                    10
                                </Dropdown.Item>
                                <Dropdown.Item onClick={() => this.onDropdownClick(25)}>
                                    25
                                </Dropdown.Item>
                                <Dropdown.Item onClick={() => this.onDropdownClick(50)}>
                                    50
                                </Dropdown.Item>
                                <Dropdown.Item onClick={() => this.onDropdownClick(100)}>
                                    100
                                </Dropdown.Item>
                            </Dropdown>
                        </div>
                        <div className="pl-5">
                            {!isLoading && (
                                <button onClick={this.onCreateClick} className="aspect-square bg-gradient-to-r from-rose-400 to-pink-700 rounded-full w-10">
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
                        <TableHeadCell>Name</TableHeadCell>
                        <TableHeadCell>Select</TableHeadCell>
                        <TableHeadCell>Delete</TableHeadCell>
                    </TableHead>
                    <TableBody className="bg-neutral-200">
                        {this.state.Files >= 0 && ([0,0,0,0,0,0,0,0,0,0]).map((i) => {
                            return (
                                <TableRow className="animate-pulse">
                                    <TableCell><div className="h-2 bg-slate-500 rounded-full w-2"></div></TableCell>
                                    <TableCell><div className="h-2 bg-slate-500 rounded"></div></TableCell>
                                </TableRow>
                            )
                        })}
                        {tableList != null && tableList.length > 0 && tableList.map((bucket: Bucket, index) => {
                            if((index >= ((this.state.CurrentPage-1) * this.state.ItemsPerPage)) &&
                                index < this.state.CurrentPage * this.state.ItemsPerPage){
                                let icon = bucket.name === Bucket ? faCircleCheckSolid : faCircleCheckRegular
                                return (
                                    <TableRow>
                                        <TableCell>{bucket.name}</TableCell>
                                        <TableCell align={"right"} width={"3rem"}><Button onClick={() => {changeBucketHandler(bucket.name)}} className="border-0 bg-gradient-to-r from-blue-400 to-cyan-700"><FontAwesomeIcon icon={icon} /></Button></TableCell>
                                        <TableCell align={"right"} width={"3rem"}><Button onClick={() => {this.deleteBucket(bucket.name)}} className="border-0 bg-gradient-to-r from-red-400 to-rose-800"><FontAwesomeIcon icon={faTrash} /></Button></TableCell>
                                    </TableRow>
                                )}
                        })}
                    </TableBody>
                </Table>
            </div>
        );



    }
})