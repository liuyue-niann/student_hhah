// 模拟文件对象
function FileInfo(name, size, status) {
    this.name = name;
    this.size = size;
    this.status = status;
}

const addFileBtn = document.getElementById('add-file-btn');
const addFolderBtn = document.getElementById('add-folder-btn');
const submit = document.getElementById('submit');
const stopBtn = document.getElementById('stop-btn');
const fileList = document.getElementById('file-list');
const fileCountInfo = document.getElementById('file-count-info');
const sizeInfo = document.getElementById('size-info');
const statusInfo = document.getElementById('status-info');
const progressBar = document.getElementById('progress-bar');
const file_upload = document.getElementById('file-upload')
let files = [];

// 添加文件模拟函数
addFileBtn.addEventListener('click', () => {
    file_upload.click()
});
submit.addEventListener('click', function () {
    file_upload.click()
})





// 监听文件选择输入框的 change 事件
file_upload.addEventListener('change', function () {

        // 获取用户选择的文件
    const file = this.files[0];
    if (file) {
        console.log('选择的文件：', file.name);
        // 获取文件扩展名
        const fileExtension = file.name;

        // 生成时间戳
        const timestamp = new Date().getTime();

        // 生成新的文件名，格式为：时间戳.扩展名
        const newFileName = `${timestamp}-${fileExtension}`;

        // 这里可以添加上传文件到服务器的代码
        // 例如使用 FormData 和 fetch API
        const formData = new FormData();
        formData.append('file', file,newFileName);
        fileCountInfo.innerText="已提交"
        fileCountInfo.style.backgroundColor= 'rgb(243, 244, 248)'
        // 模拟上传请求
        fetch('http://localhost:7788/upload', {
            method: 'POST',
            body: formData
        })
            .then(response => response.json())
            .then(data => {

            })
            .catch(error => console.error('上传失败：', error));
    }});