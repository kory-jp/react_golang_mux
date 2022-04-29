import { Input } from "@mui/material";
import { VFC } from "react";

type Props = {
  placeholder: string
  type: string
  value: string
  required: boolean
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void
}

const AuthInput: VFC<Props> = (props) => {
  return (
    <>
      <Input 
        placeholder={props.placeholder}
        type={props.type}
        value={props.value}
        required={props.required}
        onChange={props.onChange}
        sx={{
          color: "#FFF",
          backgroundColor: '#605D5D',
          borderRadius: '10px',
          height: {
            xs: '40px',
          },
          width: '100%',
          fontFamily: 'Noto Serif JP, serif',
        }}
      />
    </>
  )
}

export default AuthInput;