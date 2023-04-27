import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Component } from "react";
export class SidebarIcon extends Component {
    render() {
        const { onClick, icon } = this.props;
        return (
            <button style={{width: "40px"}} onClick={onClick} className="w-px-24 drop-shadow-2xl aspect-square rounded p-2 bg-gradient-to-r from-green-400 to-blue-500 hover:from-pink-500 hover:to-yellow-500">
                <FontAwesomeIcon icon={icon} />
            </button>
        );
    }
}

export class SidebarUser extends Component {
    render() {
        const { auth0 } = this.props;
        const { user, logout } = auth0;
        return (
            <div style={{backgroundImage: user.picture}} onClick={() => logout({ returnTo: window.location.origin })} className="drop-shadow-2xl aspect-square rounded">
                <img src={user.picture} className="rounded" style={{width: "40px"}} alt="" />
            </div>
        
        );
    }
}