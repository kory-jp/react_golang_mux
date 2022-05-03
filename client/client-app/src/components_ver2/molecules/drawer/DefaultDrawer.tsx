import { Button, Divider, Drawer } from "@mui/material";
import { Box } from "@mui/system";
import { FC } from "react";
import DriveFileRenameOutlineIcon from '@mui/icons-material/DriveFileRenameOutline';
import LogoutIcon from '@mui/icons-material/Logout';
import { User } from "../../../reducks/users/types";

type Props = {
  open: boolean
  user: User
  onClickCloseDrawer: () => void
  onClickToNewTodo: () => void
  onClickLogout: () => void
}

export const DefautlDrawer: FC<Props> = (props) => {

  const {open, user, onClickCloseDrawer, onClickToNewTodo, onClickLogout} = props
  
  return (
    <>
      <Drawer
        anchor="left"
        open={open}
        onClose={onClickCloseDrawer}
      >
        <Box
          width="200px"
        >
          <Box
            textAlign="center"
            marginY="20px"
            fontWeight="bold"
          >
            <p>
              {user.name}
            </p>
          </Box>
          <Divider />
          <Box
            textAlign="center"
            marginY="15px"
          >
            <Box
              marginBottom="10px"
            >
              <Button
                sx={{
                  color: '#424242',
                  fontFamily: 'Noto Serif JP, serif',
                }}
                onClick={onClickToNewTodo}
              >
                <DriveFileRenameOutlineIcon />
                タスク追加
              </Button>
            </Box>
            <Box>
              <Button
                sx={{
                  color: '#424242',
                  fontFamily: 'Noto Serif JP, serif',
                }}
                onClick={onClickLogout}
              >
                <LogoutIcon />
                ログアウト
              </Button>
            </Box>
          </Box>
        </Box>
      </Drawer>
    </>
  )
}

export default DefautlDrawer;