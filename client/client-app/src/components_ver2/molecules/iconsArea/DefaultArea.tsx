import { Button, Checkbox, FormControlLabel, Grid } from "@mui/material";
import { Box } from "@mui/system";
import { FC, ReactNode } from "react";
import DeleteIcon from '@mui/icons-material/Delete';

type Props = {
  children: ReactNode,
  finish: boolean,
  onChangeIsFinished: () => void,
  onClickDelete: () => void,
  onClickSomeEvent: () => void,
}

export const DefaultArea: FC<Props> = (props) => {
  const { children, finish, onChangeIsFinished, onClickDelete, onClickSomeEvent} = props

  return(
    <>
      <Grid
        container
        justifyContent='space-between'
      >
        <Grid item xs={5}>
          <FormControlLabel
            control={<Checkbox 
                        checked={finish}
                        value={finish}
                        onChange={onChangeIsFinished}
                        sx={{
                          color: '#587FBA',
                        }}
                      />} 
            sx={{
              color: '#FFF',
              fontSize: '8px',
              marginBottom: '3px',
              fontFamily: 'Merriweather, serif',
            }}
            label="FINISH"
          />
        </Grid>
        <Grid 
          item 
          xs={7}
          sx={{
            display: 'flex',
            justifyContent: 'flex-end'
          }}
        >
          <Box
            marginRight='8px'
          >
            <Button
              onClick={onClickDelete}
              sx={{
                color: '#FFF',
                fontSize: '8px',
                marginBottom: '3px',
                fontFamily: 'Merriweather, serif',
              }}
            >
              <DeleteIcon 
                sx={{
                  color: '#587FBA',
                }}
              />
              Delete
            </Button>
          </Box>
          <Box>
            <Button
              onClick={onClickSomeEvent}
              sx={{
                color: '#FFF',
                fontSize: '8px',
                fontFamily: 'Merriweather, serif',
                paddingRight: '0',
              }}
            >
              {children}
            </Button>
          </Box>
        </Grid>
      </Grid>
    </>
  )
}

export default DefaultArea;