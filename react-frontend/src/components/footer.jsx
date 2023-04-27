import {Component} from "react";
import {Footer} from "flowbite-react";

export class FooterComponent extends Component {
    render() {
        return(
            <div className="flex-row">
                <Footer container={true}>
                    <Footer.Copyright
                        href="#"
                        by="Flowbite™"
                        year={2022}
                    />
                    <Footer.LinkGroup>
                        <Footer.Link href="#">
                            About
                        </Footer.Link>
                        <Footer.Link href="#">
                            Privacy Policy
                        </Footer.Link>
                        <Footer.Link href="#">
                            Licensing
                        </Footer.Link>
                        <Footer.Link href="#">
                            Contact
                        </Footer.Link>
                    </Footer.LinkGroup>
                </Footer>
            </div>
        )
    }
}