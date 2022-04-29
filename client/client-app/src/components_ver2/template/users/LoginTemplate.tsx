import { FC } from "react";
import LoginSection from "../../organisms/users/LoginSection";
import { DefaultTemplate } from "./DefaultTemplate";

export const LoginTempate: FC = () => {
  return(
    <>
      <DefaultTemplate>
        <LoginSection />
      </DefaultTemplate>
    </>
  )
}

export default LoginTempate;