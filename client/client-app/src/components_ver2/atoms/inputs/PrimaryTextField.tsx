import { TextField } from "@mui/material";
import { FC } from "react";
import { MultiValueGeneric } from "react-select/dist/declarations/src/components/MultiValue";

type Props = {
  label: string,
  type: string,
  value: string
  mulitline: boolean,
  rows?: number,
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void,
}

export const PrimaryTextField: FC<Props> = (props) => {
  const {label, type, value, mulitline, rows, onChange} = props
  if (mulitline) {
    return(
      <>
        <TextField
          variant="standard" 
          label={label}
          fullWidth
          multiline
          type={type}
          value={value}
          rows={rows}
          onChange={onChange}
          InputLabelProps={{
            style: {
              color: '#FFF',
              fontFamily: 'Noto Serif JP, serif',
            }
          }}
          sx={{
            '& .MuiInputBase-input': {
              color: '#FFF',
              fontFamily: 'Noto Serif JP, serif',
            },
            '& .MuiInput-underline:before': {
              borderBottomColor: '#FFF'
            }
          }}
        />         
      </>
    ) 
  } else {
    return(
      <>
        <TextField
          variant="standard" 
          label={label}
          fullWidth
          type={type}
          value={value}
          onChange={onChange}
          InputLabelProps={{
            style: {
              color: '#FFF',
              fontFamily: 'Noto Serif JP, serif',
            }
          }}
          sx={{
            '& .MuiInputBase-input': {
              color: '#FFF',
              fontFamily: 'Noto Serif JP, serif',
            },
            '& .MuiInput-underline:before': {
              borderBottomColor: '#FFF'
            }
          }}
        />         
      </>
    ) 
  }
}

export default PrimaryTextField;