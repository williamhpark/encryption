import React from "react";
import "./Header.css";

import { Method } from "../../App";

type HeaderProps = {
  method: Method;
};

const Header: React.FunctionComponent<HeaderProps> = ({ method }) => {
  const descriptions: { [key: string]: string } = {
    [Method.Caesar]:
      "A Caesar cipher is a simple method of encoding messages. Caesar ciphers use a substitution method where letters in the alphabet are shifted by some fixed number of spaces to yield an encoding alphabet. A Caesar cipher with a shift of 1 would encode an A as a B, an M as an N, and a Z as an A, and so on.",
    [Method.AES]:
      "The Advanced Encryption Standard (AES) is a fast and secure form of encryption that keeps prying eyes away from our data. We see it in messaging apps like WhatsApp and Signal, programs like VeraCrypt and WinZip, in a range of hardware and a variety of other technologies that we use all of the time.",
    [Method.RSA]:
      "Under RSA encryption, messages are encrypted with a code called a public key, which can be shared openly. Due to some distinct mathematical properties of the RSA algorithm, once a message has been encrypted with the public key, it can only be decrypted by another key, known as the private key. Each RSA user has a key pair consisting of their public and private keys. As the name suggests, the private key must be kept secret.",
  };

  return (
    <div className="header">
      <h1>{method}</h1>
      <p>{descriptions[method]}</p>
    </div>
  );
};

export default Header;
