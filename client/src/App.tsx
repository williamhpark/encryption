import React, { useState } from "react";
import "./App.css";

import Dropdown from "./components/Dropdown/Dropdown";
import Header from "./components/Header/Header";

export enum Method {
  Caesar = "Caesar Cipher",
  AES = "Advanced Encryption Standard (AES)",
  RSA = "Advanced Encryption Standard (RSA)",
}

const methods = Object.values(Method);

const App = () => {
  const [method, setMethod] = useState(Method.Caesar);

  return (
    <div className="App">
      <Header method={method} />
      <Dropdown methods={methods} method={method} setMethod={setMethod} />
      {method === Method.Caesar ? (
        <p>{Method.Caesar}</p>
      ) : method === Method.AES ? (
        <p>{Method.AES}</p>
      ) : method === Method.RSA ? (
        <p>{Method.RSA}</p>
      ) : null}
    </div>
  );
};

export default App;
