import React, { Dispatch, SetStateAction } from "react";
import Box from "@mui/material/Box";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import FormControl from "@mui/material/FormControl";
import Select, { SelectChangeEvent } from "@mui/material/Select";
import "./Dropdown.css";

import { Method } from "../../App";

type DropdownProps = {
  methods: Method[];
  method: Method;
  setMethod: Dispatch<SetStateAction<Method>>;
};

const Dropdown: React.FunctionComponent<DropdownProps> = ({
  methods,
  method,
  setMethod,
}) => {
  const handleChange = (event: SelectChangeEvent) => {
    setMethod(event.target.value as Method);
  };

  return (
    <Box className="select">
      <FormControl fullWidth>
        <InputLabel id="method-select-label">Encrpytion Method</InputLabel>
        <Select
          labelId="method-select-label"
          id="method-select"
          value={method}
          label="Encryption Method"
          onChange={handleChange}
        >
          {methods.map((method) => (
            <MenuItem key={method} value={method}>
              {method}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </Box>
  );
};

export default Dropdown;
