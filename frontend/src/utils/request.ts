import axios from "axios";

import { getBaseUrl } from "./envUrl";

import router from "@/router/index";

const request = axios.create({
  baseURL: getBaseUrl(),
  timeout: 8000, // request timeout
});

request.interceptors.request.use(
  (config) => {
    if (config.method === "post" || config.method === "put") {
    }
    // console.log("request.interceptors.request");
    // do something before request is sent
    config.headers = {
      access_token: localStorage.getItem("access_token") || "",
      ...config.headers,
    };
    // config.headers["access_token"] = localStorage.getItem("access_token") | "";
    return config;
  },
  (error) => {
    console.log("request.interceptors.request error");
    return Promise.reject(error);
  }
);
request.interceptors.response.use(
  (response) => {
    // console.log("request.interceptors.response ");
    const { data } = response;
    // console.log(data);
    if (data.code && data.code > 500 && data.code < 510) {
      localStorage.setItem("access_token", "");
      window.location.reload();
      // router.back();
    }
    if (data.code == 10000) {
      return data;
    } else {
      if (data && data.msg) {
        return Promise.reject(data.msg);
      } else {
        return Promise.reject(data.toString());
      }
    }
  },
  (error) => {
    // console.log("request.interceptors.response error");
    return Promise.reject(error.toString());
  }
);

export async function wraprequest(payload: any) {
  const { umimethod, umipath, ...rest } = payload;

  if (umimethod === "post" || umimethod === "put" || umimethod === "POST" || umimethod === "PUT") {
    return request({
      url: umipath,
      method: umimethod,
      data: { ...rest },
    });
  } else {
    return request({
      url: umipath,
      method: umimethod,
      params: { ...rest },
    });
  }
}

export default request;
