function mypost(api, jsondata){
    return new Promise(res=>{
        fetch("http://localhost/api"+api,{
            method:'POST',
            body:JSON.stringify(jsondata),
            mode: 'cors',
            credentials: 'include',
        })  
        .then(response => response.json())
        .then(data => {
            if(data.status===false){alert(data.msg)}
            console.log("后端返回 ",data)
            res(data)
        });
    })
}

// ############################################
// # api_unauth.go
mypost("/unauth/account/login",{
    name:"admin", pass:"dingxin"
})
{
    "data": {
        "key": 1,
        "org": "管理员",
        "name": "admin",
        "pass": "dingxin"
    },
    "status": true
}
// ############################################
// # api_account.go
mypost("/account/logout")
{status: true}

// -----------------------------------
mypost("/account/change_password",{
    passold:"dingxin",passnew:"dingxin2"
})
{msg: '修改密码成功', status: true}

// -----------------------------------
mypost("/account/create",{
    name:"pyxfgj",
    pass:"123456",
    org:"平邑县发改局"
})
data: {key: 2, org: '平邑县发改局', name: 'pyxfgj', pass: '123456'}
msg: "创建账号成功"
status: true

// -----------------------------------
mypost("/account/edit",{
    key:2,
    name:"pyxfgj",
    pass:"123456",
    org:"平邑县3发改局"
})
data:{key: 2, org: '平邑县3发改局', name: 'pyxfgj', pass: '123456'}
msg:"handler_api_account_edit:编辑账号成功"
status:true

// -----------------------------------
mypost("/account/delete",{
    key:2
})
{msg: '删除账号成功', status: true}

// -----------------------------------
mypost("/account/getall")
data: (2) [{…}, {…}]
status: true

// -----------------------------------
mypost("/account/allkeyorg")
同上，但name&pass是空的


// ############################################
// # api_account.go