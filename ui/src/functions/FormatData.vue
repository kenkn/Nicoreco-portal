<script>
export default function formatDate(data) {
    // 引数の例: 2021-11-29T07:39:22.939Z
    // 返り値の例: 2021-11-29 07:39:22
    return data.split('.')[0].replace('T', ' ')
}
</script>

