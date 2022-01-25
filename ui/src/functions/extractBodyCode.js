import hljs from 'highlight.js';
import 'highlight.js/styles/xcode.css';
import getCodeContent from "./getCodeContent"

// extractBodyCode
// 機能 : コードハイライトの作成
export default function extractBodyCode(content) {
    let bodies = []
    let codes = []
    const re = /```[\s\S]*?```/g
    let sliceStartIdx = 0
    for (let c of content.matchAll(re)) {
        const body = content.slice(sliceStartIdx, sliceStartIdx + c.index - 1)
        const code = getCodeContent(c[0].slice(3, -3))
        const highlightedCode = '<p style="background-color: #eee">' + hljs.highlightAuto(code).value + '</p>'
        bodies.push(body)
        codes.push(highlightedCode)
        sliceStartIdx = c.index + c[0].length
    }
    bodies.push(content.slice(sliceStartIdx))

    return [bodies, codes]
}