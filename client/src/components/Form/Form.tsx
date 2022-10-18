import React, { useEffect, useState } from "react";
import axios from "axios";
import Button from "@mui/material/Button";
import TextField from "@mui/material/TextField";
import "./Form.css";

import { Method } from "../../App";

type FormProps = {
  method: Method;
  setError: React.Dispatch<React.SetStateAction<string>>;
};

const ENDPOINT = "http://localhost:8080";

const Form: React.FunctionComponent<FormProps> = ({ method, setError }) => {
  const [encryptInput, setEncryptInput] = useState("");
  const [encryptKey, setEncryptKey] = useState("");
  const [encryptMessage, setEncryptMessage] = useState("");
  const [decryptInput, setDecryptInput] = useState("");
  const [decryptKey, setDecryptKey] = useState("");
  const [decryptMessage, setDecryptMessage] = useState("");

  // If the method changes, reset all inputs and results
  useEffect(() => {
    setEncryptInput("");
    setEncryptKey("");
    setEncryptMessage("");
    setDecryptInput("");
    setDecryptKey("");
    setDecryptMessage("");
  }, [method]);

  const fetchEncryptMessage = async () => {
    setError("");

    if (!encryptInput) {
      setError('"Encrypt Message" is blank');
      return;
    }

    let methodAbbrv = "";
    let data = {};
    switch (method) {
      case Method.Caesar:
        if (!encryptKey) {
          setError('"Encrypt Key" is blank');
          return;
        }
        methodAbbrv = "caesar";
        data = {
          message: encryptInput,
          key: encryptKey,
        };
        break;
      case Method.AES:
        if (!encryptKey) {
          setError('"Encrypt Key" is blank');
          return;
        }
        methodAbbrv = "aes";
        data = {
          message: encryptInput,
          key: encryptKey,
        };
        break;
      case Method.RSA:
        methodAbbrv = "rsa";
        data = {
          message: encryptInput,
        };
    }

    await axios
      .post(`${ENDPOINT}/${methodAbbrv}/encrypt`, data)
      .then((response) => setEncryptMessage(response.data))
      .catch((error) => {
        if (error.response) {
          // The request was made and the server responded with a status code
          // that falls out of the range of 2xx
          setError(error.response.data);
        } else if (error.request) {
          // The request was made but no response was received
          // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
          // http.ClientRequest in node.js
          setError(error.request);
        } else {
          // Something happened in setting up the request that triggered an Error
          setError("Error" + error.message);
        }
      });
  };

  const fetchDecryptMessage = async () => {
    setError("");

    if (!decryptInput) {
      setError('"Decrypt Message" is blank');
      return;
    }

    let methodAbbrv = "";
    let data = {};
    switch (method) {
      case Method.Caesar:
        if (!decryptKey) {
          setError('"Decrypt Key" is blank');
          return;
        }
        methodAbbrv = "caesar";
        data = {
          message: decryptInput,
          key: decryptKey,
        };
        break;
      case Method.AES:
        if (!decryptKey) {
          setError('"Decrypt Key" is blank');
          return;
        }
        methodAbbrv = "aes";
        data = {
          message: decryptInput,
          key: decryptKey,
        };
        break;
      case Method.RSA:
        methodAbbrv = "rsa";
        data = {
          message: decryptInput,
        };
    }

    await axios
      .post(`${ENDPOINT}/${methodAbbrv}/decrypt`, data)
      .then((response) => setDecryptMessage(response.data))
      .catch((error) => {
        if (error.response) {
          // The request was made and the server responded with a status code
          // that falls out of the range of 2xx
          setError(error.response.data);
        } else if (error.request) {
          // The request was made but no response was received
          // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
          // http.ClientRequest in node.js
          setError(error.request);
        } else {
          // Something happened in setting up the request that triggered an Error
          setError("Error" + error.message);
        }
      });
  };

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
          <Button
            className="submit-btn"
            variant="contained"
            onClick={fetchEncryptMessage}
          >
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
          <Button
            className="submit-btn"
            variant="contained"
            onClick={fetchDecryptMessage}
          >
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
