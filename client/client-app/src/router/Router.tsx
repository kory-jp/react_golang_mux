import { VFC } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Registration from "../components/pages/auth/Registration";

export const Router: VFC = () => {
  return(
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Registration/>}/>
      </Routes>
    </BrowserRouter>
  )
}

export default Router;
