export default function getCodeContent(code) {
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