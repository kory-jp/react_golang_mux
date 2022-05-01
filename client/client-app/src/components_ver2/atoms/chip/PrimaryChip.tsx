import { Chip } from "@mui/material";
import { FC } from "react";
type Props = {
  label: string,
  onClick?: () => void,
  colorFlag?: number | undefined,
}

export const PrimaryChip: FC<Props> = (props) => {
  const {label, onClick, colorFlag} = props
  return(
    <Chip
    label={label}
    onClick={onClick}
    sx={{
      color: '#FFF',
      backgroundColor: colorFlag === undefined || colorFlag === 1 ? '#587FBA' : '#605D5D',
      height: '24px',
      width: '56px',
      fontSize: '8px',
    }}
  />
  )
}