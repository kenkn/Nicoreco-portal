import hljs from 'highlight.js';
import 'highlight.js/styles/xcode.css';
import getCodeContent from "./getCodeContent"

// extractBodyCode
// 機能 : コードハイライト用の
// HACK 
//   * エスケープ文字によるコード回避，正規表現によるコード部分抽出
// "```" でbodyを分割すると，偶数idxの時は非コード部分，奇数idxの時はコード部分と考えることができる．
export default function extractBodyCode(content) {
    let bodies = []
    let codes = []
    const splitedContent = content.split('```')
    for (let i = 0; i < splitedContent.length; ++i) {
        if (i % 2 == 1) {
            const code = getCodeContent(splitedContent[i])
            const highlightedCode = '<p style="background-color: #eee">' + hljs.highlightAuto(code).value + '</p>'
            codes.push(highlightedCode)
        } else {
            bodies.push(getCodeContent(splitedContent[i]))
        }
    }

    return [bodies, codes]
}