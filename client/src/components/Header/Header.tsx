import React from "react";
import "./Header.css";

import { Method } from "../../App";

type HeaderProps = {
  method: Method;
};

const Header: React.FunctionComponent<HeaderProps> = ({ method }) => {
  const descriptions: { [key: string]: string } = {
    [Method.Caesar]: "Caesar",
    [Method.AES]: "AES",
    [Method.RSA]: "RSA",
  };

  return (
    <div className="header">
      <h1>{method}</h1>
      <p>{descriptions[method]}</p>
    </div>
  );
};

export default Header;
