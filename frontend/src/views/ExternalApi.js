import React, { useState } from "react";
import { Button, Alert } from "reactstrap";
import Highlight from "../components/Highlight";
import { useAuth0, withAuthenticationRequired } from "@auth0/auth0-react";
import { getConfig } from "../config";
import Loading from "../components/Loading";
import * as CyanAPI from "ts-cloudwpss23-openapi-cyan";
import * as $ from 'jquery';
import {User} from "ts-cloudwpss23-openapi-cyan";
import { useFilePicker } from 'use-file-picker';

export const ExternalApiComponent = () => {
  const { apiOrigin = "https://api.cloudwp.anwski.de", audience } = getConfig();
  const api = new CyanAPI.DefaultApi("https://openapi-asog6d6nbq-ez.a.run.app");
  const [openFileSelector, { filesContent, loading, plainFiles }] = useFilePicker({
    accept: 'image/*'
  });

  const [state, setState] = useState({
    showResult: false,
    apiMessage: "",
    error: null,
    showFiles: false,
    bucket: "",
    files: []
  });

  const {
    getAccessTokenSilently,
    loginWithPopup,
    getAccessTokenWithPopup,
  } = useAuth0();

  const handleConsent = async () => {
    try {
      await getAccessTokenWithPopup();
      setState({
        ...state,
        error: null,
      });
    } catch (error) {
      setState({
        ...state,
        error: error.error,
      });
    }

    await callApi();
  };

  const handleLoginAgain = async () => {
    try {
      await loginWithPopup();
      setState({
        ...state,
        error: null,
      });
    } catch (error) {
      setState({
        ...state,
        error: error.error,
      });
    }

    await callApi();
  };

  const callApi = async () => {
    try {
      let prom = api.getUsersUserId(100);
      $.when(prom).then(function (status){
        let body: User = status.body;
        console.log(body.name);
        setState({
          ...state,
          showResult: true,
          apiMessage: body,
        });
      })
    } catch (error) {
      setState({
        ...state,
        error: error.error,
      });
    }
  };

  const handle = (e, fn) => {
    e.preventDefault();
    fn();
  };

  getAccessTokenSilently().then(function (token){
    api.configuration.accessToken = token
    if(!state.showFiles) {
      $.when(api.getFiles()).then(function (status) {
        let response: CyanAPI.GetFiles200Response = status.body;
        setState({
          ...state,
          showFiles: true,
          bucket: response.bucket,
          files: response.files
        });
      });
    }
  })
  return (
    <>
      <div className="mb-5">
        {state.error === "consent_required" && (
          <Alert color="warning">
            You need to{" "}
            <a
              href="#/"
              class="alert-link"
              onClick={(e) => handle(e, handleConsent)}
            >
              consent to get access to users api
            </a>
          </Alert>
        )}

        {state.error === "login_required" && (
          <Alert color="warning">
            You need to{" "}
            <a
              href="#/"
              class="alert-link"
              onClick={(e) => handle(e, handleLoginAgain)}
            >
              log in again
            </a>
          </Alert>
        )}



        <h1>External API</h1>
        <p className="lead">
          Ping an external API by clicking the button below.
        </p>

        {state.showFiles && (
            <div>
              <h2>External API</h2>
              <p className="lead">
                List of { state.bucket } Bucket files
              </p>
              {state.files.map(function (name) {
                  return <p>{name}</p>
              })}
            </div>
        )}

        <p>
          This will call a local API on port 3001 that would have been started
          if you run <code>npm run dev</code>. An access token is sent as part
          of the request's `Authorization` header and the API will validate it
          using the API's audience value.
        </p>

        {!audience && (
          <Alert color="warning">
            <p>
              You can't call the API at the moment because your application does
              not have any configuration for <code>audience</code>, or it is
              using the default value of <code>YOUR_API_IDENTIFIER</code>. You
              might get this default value if you used the "Download Sample"
              feature of{" "}
              <a href="https://auth0.com/docs/quickstart/spa/react">
                the quickstart guide
              </a>
              , but have not set an API up in your Auth0 Tenant. You can find
              out more information on{" "}
              <a href="https://auth0.com/docs/api">setting up APIs</a> in the
              Auth0 Docs.
            </p>
            <p>
              The audience is the identifier of the API that you want to call
              (see{" "}
              <a href="https://auth0.com/docs/get-started/dashboard/tenant-settings#api-authorization-settings">
                API Authorization Settings
              </a>{" "}
              for more info).
            </p>

            <p>
              In this sample, you can configure the audience in a couple of
              ways:
            </p>
            <ul>
              <li>
                in the <code>src/index.js</code> file
              </li>
              <li>
                by specifying it in the <code>auth_config.json</code> file (see
                the <code>auth_config.json.example</code> file for an example of
                where it should go)
              </li>
            </ul>
            <p>
              Once you have configured the value for <code>audience</code>,
              please restart the app and try to use the "Ping API" button below.
            </p>
          </Alert>
        )}

        <Button
          color="primary"
          className="mt-5"
          onClick={callApi}
          disabled={!audience}
        >
          Ping API
        </Button>
        <br/>
        <Button
            color="primary"
            className="mt-5"
            onClick={() => {
              try {
                const result = openFileSelector();
              } catch (err) {
                console.log(err);
                console.log('Something went wrong or validation failed');
              }
            }}
            disabled={!audience}
        >
          Select File
        </Button>
        <br />
        {plainFiles.map((file, index) => (
            <div>
              <h2>{file.name}</h2>
              <Button
                  color="primary"
                  className="mt-5"
                  onClick={() => {
                    console.log(file.size);
                    $.when(api.putFileUpload(file.name, file)).then(function (status){
                      console.log(status)
                    })
                  }}
                  disabled={!audience}
              >
                Upload File
              </Button>
              <br />
            </div>
        ))}
      </div>

      <div className="result-block-container">
        {state.showResult && (
          <div className="result-block" data-testid="api-result">
            <h6 className="muted">Result</h6>
            <Highlight>
              <span>{JSON.stringify(state.apiMessage, null, 2)}</span>
            </Highlight>
          </div>
        )}
      </div>
    </>
  );
};

export default withAuthenticationRequired(ExternalApiComponent, {
  onRedirecting: () => <Loading />,
});
