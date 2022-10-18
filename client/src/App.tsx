import React, { useState } from "react";
import "./App.css";

import Dropdown from "./components/Dropdown/Dropdown";
import Form from "./components/Form/Form";
import Header from "./components/Header/Header";

export enum Method {
  Caesar = "Caesar Cipher",
  AES = "Advanced Encryption Standard (AES)",
  RSA = "Advanced Encryption Standard (RSA)",
}

const methods = Object.values(Method);

const App = () => {
  const [method, setMethod] = useState(Method.Caesar);
  const [error, setError] = useState("");

  return (
    <div className="App">
      <Header method={method} />
      <Dropdown methods={methods} method={method} setMethod={setMethod} />
      <Form method={method} setError={setError} />
      {error ? <h3 className="error-msg">{error}</h3> : null}
    </div>
  );
};

export default App;
