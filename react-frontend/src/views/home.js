import { Component } from "react";
import {Alert, Button} from "flowbite-react";
import {HiInformationCircle} from "react-icons/hi";
import * as $ from "jquery";
import {apiClient} from "../api";
import {PostV1BucketNameRequest} from "ts-cloudwpss23-openapi-cyan";

export class PageHome extends Component {

    constructor(props) {
        super(props);
    }


    render() {
        return(
            <div className="p-10 w-full">
                <Alert
                    color="failure"
                    icon={HiInformationCircle}
                >
                  <span>
                    <span className="font-medium">
                      Info alert!
                    </span>
                      {' '}Please Login to use this Service!

                  </span>

                </Alert>
            </div>
        )
    }
}