import { Input, styled } from "@mui/material"
import { FC } from "react"

type Props = {
  placeholder: string
  type: string
  value: string
  required: boolean
  onChange: (event: React.ChangeEvent<HTMLInputElement>) => void
}

const PrimaryInputStyle = styled(Input)((props) => ({
  marginRight: "10px",
  placeholder: props.placeholder,
}))

export const PrimaryInput: FC<Props> = (props: Props) => {
  return(
    <PrimaryInputStyle 
      placeholder={props.placeholder}
      type={props.type}
      value={props.value}
      required={props.required}
      onChange={props.onChange}
    />
  )
}