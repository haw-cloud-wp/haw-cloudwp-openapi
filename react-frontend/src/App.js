import './App.css';
import { Sidebar } from './components/sidebar';

import { PageHome } from './views/home';
import {Routes, Route, BrowserRouter, Router} from 'react-router-dom';
import { useAuth0 } from '@auth0/auth0-react';
import {auth0, CallQue, SetAuth0} from "./auth0";
import {apiClient} from "./api";
import {ViewBucketList} from "./views/filelist";
import {getConfig} from "./config";
function App() {
    SetAuth0(useAuth0());
    const config = getConfig();
    if(auth0.isAuthenticated){
        console.log(auth0.user)
        const tokenOptions = {
            authorizationParams: {
                audience: config.audience,
                scope: "read:bucket_customer_bucker BucketA BucketB BucketC"
            }
        };
        auth0.getAccessTokenSilently(tokenOptions).then(function (token) {
                apiClient.configuration.accessToken = token;
                CallQue();
            }).catch((e) => {
            auth0.getAccessTokenWithPopup(tokenOptions).then(function (token) {
                apiClient.configuration.accessToken = token;
                CallQue();
            }).catch((e) => {
                console.log("Token error! : ", e)
            })
        })
    }
    return (

        <div className="flex-col pl-16 w-full">
            <Sidebar />
            <Routes>
                <Route path="/">
                    <Route index={true} element={<PageHome />} />
                    <Route path="/objectlist" element={<ViewBucketList />} />
                </Route>
            </Routes>
        </div>
  );
}

export default App;
