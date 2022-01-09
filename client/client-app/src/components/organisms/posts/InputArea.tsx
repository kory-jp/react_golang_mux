import { Box, Input } from "@mui/material"
import { PrimaryButton } from "../../atoms/button/PrimaryButton"
import { PrimaryInput } from "../../atoms/input/PrimaryInput"

export const InputArea = () => {
  return(
    <Box>
      <PrimaryInput placeholder={'入力'}/>
      <PrimaryButton>
        送信
      </PrimaryButton>
    </Box>
  )
}

export default InputArea;