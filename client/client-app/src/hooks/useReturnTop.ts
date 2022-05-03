export const useReturnTop = () => {
  const returnTop = () => {
    window.scrollTo({
      top: 0,
      behavior: 'smooth'
    })
  }

  return returnTop;
}

export default useReturnTop;