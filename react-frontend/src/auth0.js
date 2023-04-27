import {Auth0ContextInterface} from "@auth0/auth0-react";
import {type} from "@testing-library/user-event/dist/type";

export let auth0: Auth0ContextInterface;

export function SetAuth0(at: Auth0ContextInterface){
    auth0 = at;
}

let promises : Function[] = [];
export function Auth0AddQue(call: Function){
    if(auth0.isAuthenticated){
        call();
    }
    promises.push(call);
}

export function CallQue(){
    promises.forEach(function (call){
        call();
    })
}