import { useSelector } from "react-redux"
import { RooState } from "../reducks/store/store";

export const useLoadingState = () => {
  const loadingState = useSelector((state: RooState) => state.loading.status);
  return loadingState
}

export default useLoadingState;