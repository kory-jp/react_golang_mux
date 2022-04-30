import { FC } from "react";
import RegistrationSection from "../../organisms/users/RegistrationSection";
import { DefaultTemplate } from "./DefaultTemplate";

export const RegistrationTemplate: FC = () => {
  return(
    <DefaultTemplate>
      <RegistrationSection />
    </DefaultTemplate>
  )
}

export default RegistrationTemplate;