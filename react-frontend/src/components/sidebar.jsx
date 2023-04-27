import React, { Component } from "react";
import { faBurger, faFile, faRightToBracket, faGear, faMoon, faSun } from "@fortawesome/free-solid-svg-icons";
import { SidebarIcon, SidebarUser } from "./sidebar/icons";
import { SidebarSubmenuFiles } from "./sidebar/submenuFiles";
import {auth0} from "../auth0";
import {Link} from "react-router-dom";
import {toggleTheme, updateTheme, isDark} from "../utils/theme";
import {withAuth0} from "@auth0/auth0-react";

export const Sidebar = withAuth0(class Sidebar extends Component{
    constructor(params){
        super(params);
        this.state = {
            open: true,
            dark: isDark()
        }
    }
    render() {
        return(
            <div className="fixed top-0 left-0">
                <div className="flex flex-col items-center w-16 h-screen py-8 space-y-8 bg-white dark:bg-gray-900 dark:border-gray-700 shadow">
                    <Link to="/"><SidebarIcon onClick={() => {this.setState({open: !this.state.open});}} icon={faBurger} /></Link>
                    { this.props.auth0.isAuthenticated && (<Link to="/objectlist"><SidebarIcon icon={faFile} /></Link>)}
                    { this.props.auth0.isAuthenticated && (<SidebarIcon icon={faGear} />)}
                    <div className="grow" />
                    <SidebarIcon onClick={()=>{toggleTheme();this.setState({dark:isDark()});}} icon={ this.state.dark ? faSun: faMoon } />
                    { !this.props.auth0.isAuthenticated && ( <SidebarIcon onClick={() => this.props.auth0.loginWithRedirect({})} icon={faRightToBracket} /> )}
                    { this.props.auth0.isAuthenticated && (<SidebarUser auth0={auth0}/>)}
                </div>
            </div>
        );
    }
})
