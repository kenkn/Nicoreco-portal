import hljs from 'highlight.js';
import 'highlight.js/styles/xcode.css';

export default function extractHighlightPart(bodyStr) {
    let contents = []
    let isMathjaxStarted
    let isCodeStarted
    let textStartIdx = 0
    let mathjaxStartIdx
    let codeStartIdx
    let content

    const initStarted = () => {
        isMathjaxStarted = false
        isCodeStarted = false
    }

    initStarted()

    for (let i = 0; i < (bodyStr.length-2); ++i) {
        if (bodyStr.slice(i, i+2) === "$$") {
            // mathjax pattern
            // ex.) '$$x = {-b \pm \sqrt{b^2-4ac} \over 2a}.$$'
            if (isCodeStarted === true) {
                continue
            }
            if (isMathjaxStarted === false) {
                isMathjaxStarted = true
                mathjaxStartIdx = i
            } else {
                initStarted()
                const text = bodyStr.slice(textStartIdx, mathjaxStartIdx)
                content = {
                    attr: "text",
                    body: text,
                }
                textStartIdx = i + 3
                contents.push(content)

                const math = getCodeContent(bodyStr.slice(mathjaxStartIdx, i+2))
                content = {
                    attr: "math",
                    body: math,
                }
                contents.push(content)
            }
        }
        if (bodyStr.slice(i, i+3) === '```') {
            // code pattern
            if (isMathjaxStarted === true) {
                continue
            }
            if (isCodeStarted === false) {
                isCodeStarted = true
                codeStartIdx = i
            } else {
                initStarted()
                const text = bodyStr.slice(textStartIdx, codeStartIdx)
                content = {
                    attr: "text",
                    body: text
                }
                textStartIdx = i + 3
                contents.push(content)

                const code = getCodeContent(bodyStr.slice(codeStartIdx+3, i))
                const highlightedCode = '<p style="background-color: #eee">' + hljs.highlightAuto(code).value + '</p>'
                content = {
                    attr: "code",
                    body: highlightedCode,
                }
                
                contents.push(content)
            }
        }
    }
    
    const text = bodyStr.slice(textStartIdx, bodyStr.length)
    content = {
        attr: 'text',
        body: text
    }
    contents.push(content)

    return contents
}

const getCodeContent = (code) => {
    let sliceIndex = 0
    for (let i = 0; i < code.length; ++i) {
        if (code[i] == ' ') {
            continue
        } else if (code[i] == '\n') {
            sliceIndex = i + 1
        } else {
            break
        }
    }
    return code.slice(sliceIndex)
}