let env = process.env["NODE_ENV"];

export function getBaseUrl(): string {
  return "/"
  /*
  if (env == "development") {
    return "http://localhost:1026/"; // 111
  } else {
    return "http://render.tpddns.cn:15001/"; // https
  }
  */

}

export function getActionUrl() {
  return getBaseUrl() + "/pdapi/system/upload/single/image";
}
