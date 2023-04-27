import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import { getConfig } from "./config";
import history from "./utils/history";
import {Auth0ProviderWithNavigate} from "./auth0navigate";
import {BrowserRouter} from "react-router-dom";

// On page load or when changing themes, best to add inline in `head` to avoid FOUC
if (localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches)) {
  document.documentElement.classList.add('dark')
} else {
  document.documentElement.classList.remove('dark')
}

const onRedirectCallback = (appState) => {
  history.push(
    appState && appState.returnTo ? appState.returnTo : window.location.pathname
  );
};

const config = getConfig();

const providerConfig = {
  domain: config.domain,
  clientId: config.clientId,
  onRedirectCallback,
  authorizationParams: {
    redirect_uri: window.location.origin,
    ...(config.audience ? { audience: config.audience } : null),
  },
};

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
      <BrowserRouter>
        <Auth0ProviderWithNavigate>
            <App />
        </Auth0ProviderWithNavigate>
      </BrowserRouter>
    </React.StrictMode>
);
