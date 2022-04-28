export const handleToDateFormat = (d: Date | null) => {
  if (d) {
    // 引数にstring等を受け取り、Dateオブジェクトを返却(各種メソッドをもつ)
    const date = new Date(d)
    // match+\dにて数値部分を抽出して、toJSONにてjson形式にて返却
    const j = date.toJSON().match(/\d+/g)
    if (j) {
      return `${j[0]}/${j[1]}/${j[2]} ${j[3]}:${j[4]}`
    } else {
      return d
    }
  } else {
    return d
  }
}
  
export default handleToDateFormat;