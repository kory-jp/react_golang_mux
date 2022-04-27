import { Box, Checkbox, FormControlLabel } from "@mui/material";
import { useCallback, useEffect, useState, VFC } from "react";
import { TaskCard } from "../../../reducks/taskCards/types";

type Props ={
  taskCard: TaskCard
}

export const IndexTaskCard: VFC<Props> = (props) => {
  const {taskCard} = props
  const [isFinished, setIsFinished] = useState(false)
  useEffect(() => {
    setIsFinished(taskCard.isFinished)
  },[])

  const onChangeIsFinished = useCallback(() => {
    if (isFinished) {
      setIsFinished(false)
      // dispatch(updateIsFinished(id, false))
    } else {
      setIsFinished(true)
      // dispatch(updateIsFinished(id, true))
    }
  }, [isFinished])

  return(
    <Box
      sx={{
        backgroundColor: isFinished ? 'text.disabled' : 'white',
        borderRadius: "10px",
        padding: "20px",
        marginBottom: "40px",
      }}
    >
      <Box>
        <Box
          sx={{
            fontSize: "24px"
          }}
        >
          {taskCard.title}
        </Box>
        <Box
          sx={{
            display: "flex",
            justifyContent: "space-between"
          }}
        >
          <Box>
            <FormControlLabel
              control={<Checkbox 
                          checked={isFinished}
                          value={isFinished}
                          onChange={onChangeIsFinished}
                        />} 
              label="finish"
              sx={{
                fontSize: '8px',
                marginBottom: '3px'
              }}
            />
          </Box>
          <Box
            marginY="auto"
          >
            <p>
              {taskCard.createdAt}
            </p>
          </Box>
        </Box>
      </Box>
    </Box>
  )
}

export default IndexTaskCard;