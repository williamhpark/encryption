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
      "The Advanced Encryption Standard (AES) aka Rijndael is an encryption algorithm created in 2001 by NIST. It's a symmetric block cipher. It can use a 128/192/256-bit key to encrypt (256 bit for this application), takes 128 bits as input and outputs 128 bits of encrypted cipher text as output.",
    [Method.RSA]:
      "Under RSA encryption, messages are encrypted with a code called a public key, which can be shared openly. Due to some distinct mathematical properties of the RSA algorithm, once a message has been encrypted with the public key, it can only be decrypted by another key, known as the private key (which must be kept secret). Each RSA user has a key pair consisting of their public and private keys. In this application, the public and private keys are generated and stored in the server during runtime.",
  };

  return (
    <div className="header">
      <h1>{method}</h1>
      <p>{descriptions[method]}</p>
    </div>
  );
};

export default Header;
