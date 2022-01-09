import { Input, styled } from "@mui/material"

type Props = {
  placeholder: string
}

const PrimaryInputStyle = styled(Input)((props) => ({
  marginRight: "10px",
  placeholder: props.placeholder
}))

export const PrimaryInput = (props: Props) => {
  return(
    <PrimaryInputStyle placeholder={props.placeholder}/>
  )
}
