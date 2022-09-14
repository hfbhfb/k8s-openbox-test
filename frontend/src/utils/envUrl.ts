let env = process.env["NODE_ENV"];

export function getBaseUrl(): string {
  if (env == "development") {
    return "http://localhost:1026/"; // 111
  } else {
    return "http://render.tpddns.cn:15001/"; // https
  }
  // if (REACT_APP_ENV == "development" || REACT_APP_ENV == "dev" || REACT_APP_ENV == undefined) {
  //   // return "http://10.66.66.206:20211"
  //   return "https://test-cbc.ycandyz.com";
  //   // return "https://uat-1.ycandyz.com"
  // }

  // if (REACT_APP_ENV == "test" || REACT_APP_ENV == "testing") {
  //   return "https://test-cbc.ycandyz.com";
  //   console.log = () => {};
  // }
  // if (REACT_APP_ENV == "uat") {
  //   return "https://uat-1.ycandyz.com";
  //   console.log = () => {};
  // }
  // if (REACT_APP_ENV == "product" || REACT_APP_ENV == "prod") {
  //   return "https://cbc-api.ycandyz.com";
  //   console.log = () => {};
  // }
  // return "http://localhost:9801/";
}

export function getActionUrl() {
  return getBaseUrl() + "/pdapi/system/upload/single/image";
}
