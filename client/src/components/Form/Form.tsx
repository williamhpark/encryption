import React, { useState } from "react";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import "./Form.css";

import { Method } from "../../App";

type FormProps = {
  method: Method;
};

const Form: React.FunctionComponent<FormProps> = ({ method }) => {
  const [encryptInput, setEncryptInput] = useState("");
  const [encryptKey, setEncryptKey] = useState("");
  const [encryptMessage, setEncryptMessage] = useState("");
  const [decryptInput, setDecryptInput] = useState("");
  const [decryptKey, setDecryptKey] = useState("");
  const [decryptMessage, setDecryptMessage] = useState("");

  return (
    <div className="form">
      <div className="container">
        <div className="section">
          <div className="input-container">
            <TextField
              id="outlined-basic"
              label="Encrypt Message"
              variant="outlined"
              value={encryptInput}
              onChange={(e) => setEncryptInput(e.target.value)}
            />
            {method === Method.Caesar || method === Method.AES ? (
              <TextField
                id="outlined-basic"
                label="Encrypt Key"
                variant="outlined"
                value={encryptKey}
                onChange={(e) => setEncryptKey(e.target.value)}
              />
            ) : null}
          </div>
          <Button className="submit-btn" variant="contained" onClick={() => {}}>
            Encrypt
          </Button>
        </div>
        <div className="result">
          <h3>Result</h3>
          {encryptMessage ? <p>{encryptMessage}</p> : null}
        </div>
      </div>
      <div className="container">
        <div className="section">
          <div className="input-container">
            <TextField
              id="outlined-basic"
              label="Decrypt Message"
              variant="outlined"
              value={decryptInput}
              onChange={(e) => setDecryptInput(e.target.value)}
            />
            {method === Method.Caesar || method === Method.AES ? (
              <TextField
                id="outlined-basic"
                label="Decrypt Key"
                variant="outlined"
                value={decryptKey}
                onChange={(e) => setDecryptKey(e.target.value)}
              />
            ) : null}
          </div>
          <Button className="submit-btn" variant="contained" onClick={() => {}}>
            Decrypt
          </Button>
        </div>
        <div className="result">
          <h3>Result</h3>
          {decryptMessage ? <p>{decryptMessage}</p> : null}
        </div>
      </div>
    </div>
  );
};

export default Form;
