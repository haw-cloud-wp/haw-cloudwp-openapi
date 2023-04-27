import {Auth0Provider, useAuth0} from "@auth0/auth0-react";
import React from "react";
import { useNavigate } from "react-router-dom";
import {getConfig} from "./config";

export const Auth0ProviderWithNavigate = ({ children }) => {
    const navigate = useNavigate();
    const config = getConfig();


    const onRedirectCallback = (appState) => {
        navigate(appState?.returnTo || window.location.pathname);
    };

    const providerConfig = {
        domain: config.domain,
        clientId: config.clientId,
        onRedirectCallback,
        authorizationParams: {
            redirect_uri: window.location.origin,
            ...(config.audience ? { audience: config.audience } : null),
        },
    };

    return (
        <Auth0Provider
            {...providerConfig}
        >
            {children}
        </Auth0Provider>
    );
};