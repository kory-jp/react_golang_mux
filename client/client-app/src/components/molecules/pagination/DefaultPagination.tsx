import { Pagination } from "@mui/material";
import { Box } from "@mui/system";
import { FC } from "react";


type Props = {
  count: number,
  onChange: (event: React.ChangeEvent<unknown>, page: number) => void,
  page: number
}

export const DefaultPagination: FC<Props> = (props) => {
  const {count, onChange, page} = props
  return(
    <Box>
      <Pagination 
        count={count}
        onChange={onChange}
        page={page}
        variant="outlined"
        color="primary"
      />
    </Box>
  )
}

export default DefaultPagination;