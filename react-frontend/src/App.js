import './App.css';
import { Sidebar } from './components/sidebar';
import * as React from "react"
import { PageHome } from './views/home';
import {Routes, Route, BrowserRouter, Router} from 'react-router-dom';
import {useAuth0, withAuth0} from '@auth0/auth0-react';
import {auth0, Auth0AddQue, CallQue, SetAuth0} from "./auth0";
import {apiClient} from "./api";
import {Filelist} from "./views/filelist";
import {getConfig} from "./config";
import {BucketList} from "./views/bucketlist";
import {Bucket} from "ts-cloudwpss23-openapi-cyan";
import * as $ from "jquery"
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
        <AppClass />
    );
}

const AppClass = withAuth0(class extends React.Component{
    constructor(props) {
        super(props);
        this.setBucketCallback = this.setBucketCallback.bind(this)
        this.state = {
            CurrBucket: "testbucket"
        }
        let app = this
        Auth0AddQue(() =>{
            $.when(apiClient.getV1Buckets()).then((s) => {
                let buckets : Array<Bucket> = s.body;
                app.setBucketCallback(buckets[0].name)
            })
        })
    }

    setBucketCallback(bucket){
        console.log("Set to bucket:", bucket)
        this.setState({CurrBucket: bucket})

    }
    render() {
        return (
            <div className="flex-col pl-16 w-full">
                <Sidebar />
                <Routes>
                    <Route path="/">
                        <Route index={true} element={<BucketList Bucket={this.state.CurrBucket} changeBucketHandler={this.setBucketCallback} />} />
                        <Route path="/objectlist" element={<Filelist Bucket={this.state.CurrBucket} />} />
                    </Route>
                </Routes>
            </div>
        );
    }
})

export default App;
