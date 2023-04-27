import { Component } from "react";
import {Auth0AddQue} from "../../auth0";
import * as $ from "jquery";
import {apiClient} from "../../api";
import {GetFiles200Response} from "ts-cloudwpss23-openapi-cyan";
class FileItem extends Component {
    render(){
        const { name, size, lastMod } = this.props;
        return(
            <button className="flex items-center w-full px-5 py-2 transition-colors duration-200 dark:hover:bg-gray-800 gap-x-2 hover:bg-gray-100 focus:outline-none">
                <div className="text-left rtl:text-right">
                    <h1 className="text-sm font-medium text-gray-700 capitalize dark:text-white">{name}</h1>
                    <p className="text-xs text-gray-500 dark:text-gray-400">{Math.round(size/1024)} kb</p>
                    <p className="text-xs text-gray-500 dark:text-gray-400">{lastMod}</p>
                </div>
            </button>
        );
    }
}

export class SidebarSubmenuFiles extends Component {
    
    constructor(props){
        super(props);
        this.state = {
            files: [],
            bucket: ""
        }
        let cmp = this;
        Auth0AddQue(() => {
            $.when(apiClient.getFiles()).then(function (status) {
                let response : GetFiles200Response  = status.body;
                cmp.setState({files: response.files, bucket: response.bucket})
            })
        })
    }
    render() {
        const { open } = this.props;
        return(
            <div className={`-z-20 h-screen py-8 overflow-y-auto bg-white border-l border-r sm:w-64 w-60 dark:bg-gray-900 dark:border-gray-700 transition-all ease-in-out ${open?"":"-translate-x-full"}`} >
                    <h2 className="px-5 text-lg font-medium text-gray-800 dark:text-white">{this.state.bucket}</h2>
                    <div className="mt-8 space-y-4">
                        {this.state.files.map((file) => <FileItem name={file.name} size={file.size} lastMod={file.lastmod} />)}
                    </div>
                </div>
        )
    }
}