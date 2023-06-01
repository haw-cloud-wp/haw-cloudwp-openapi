import * as API from "ts-cloudwpss23-openapi-cyan";
const localUrl = "http://localhost"
const prodUrl = "https://cloudwp-openapi.azurewebsites.net"
const isDev = false
export const apiClient = new API.DefaultApi(isDev ? localUrl : prodUrl, );