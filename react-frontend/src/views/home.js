import { Component } from "react";
import {Alert} from "flowbite-react";
import {HiInformationCircle} from "react-icons/hi";

export class PageHome extends Component {
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