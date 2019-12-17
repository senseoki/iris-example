const ADS_API_DOMAIN = "http://localhost:9001";
const $customerID = $('#customer-id');
const $excelKeyWord = $('#excel-keyWord');
const $excelAccountPerformance = $("#excel-accountPerformance");
const $excelStream = $("#excel-stream");

var excelBtnClickCheck = false;

$(document).ready(function() {
    init();
    event();
    //testLocalstorage();
});

function init() {
    //Initialize Select2 Elements
    $('.select2bs4').select2({
        theme: 'bootstrap4'
    });

    // Get CustomerId and Setting
    axios({
        method: 'GET',
        url:  `${ADS_API_DOMAIN}/account`,
    }).then(response => {
        if(response.data) {
            var cnt = 0;
            $.each(response.data[0], (k, v) => {
                ++cnt
                console.log(k, v);
                $customerID.append(`<option value="${k}">${v}</option>`);
            });

            console.log("size: ", cnt);
        }
    }).catch(function(err) {
        console.log("Get axios account err : ", err)
    });
}

function event() {

    $excelKeyWord.click((e) => {
        e.preventDefault();

        if(excelBtnClickCheck) {
            return;
        }
        excelBtnClickCheck = true;
        var customerId = $customerID.val();
        if(customerId == '') {
            alert('customerId를 선택해주세요');
            return;
        }
        var during = $('input:radio[name="duringRadio"]:checked').val();
        var fileName = makeFileName(`keyword_${customerId}`);
        excelDownAjax({
            url: `${ADS_API_DOMAIN}/report/down/keyword?customerId=${customerId}&fileName=${fileName}&during=${during}`,
            fileName: fileName
        });
    });

    $excelAccountPerformance.click((e) => {
        e.preventDefault();

        if(excelBtnClickCheck) {
            return;
        }
        excelBtnClickCheck = true;
        var customerId = $customerID.val();
        if(customerId == '') {
            alert('customerId를 선택해주세요');
            return;
        }
        var during = $('input:radio[name="duringRadio"]:checked').val();
        var fileName = makeFileName(`accountPerformance_${customerId}`);
        excelDownAjax({
            url: `${ADS_API_DOMAIN}/report/down/account?customerId=${customerId}&fileName=${fileName}&during=${during}`,
            fileName: fileName
        })
    });

    function makeFileName(downType) {
        return `${downType}_` + moment(moment.now()).format("YYYY-MM-DD_HH_mm_ss");
    }
    
    // $excelStream.click((e) => {
    //     e.preventDefault();
    //     if(excelBtnClickCheck) {
    //         return;
    //     }
    //     excelBtnClickCheck = true;
    //     var customerId = $customerID.val();
    //     if(customerId == '') {
    //         alert('customerId를 선택해주세요');
    //         return;
    //     }
    //     var during = $('input:radio[name="duringRadio"]:checked').val();
    //     var fileName = makeFileName(`keyword_${customerId}`);

    //     location.href = `${ADS_API_DOMAIN}/report/down/keyword?customerId=${customerId}&fileName=${fileName}&during=${during}`;
    //     excelBtnClickCheck = false;
    // });
}



function excelDownAjax({url, fileName}) {
    axios({
        method: 'get',
        url: url,
        responseType: 'blob',
    }).then(res => {    
        const url = window.URL.createObjectURL(new Blob([res.data], { type: "Application/Msexcel" }));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', fileName + ".xlsx");
        document.body.appendChild(link);
        link.click();    

        excelBtnClickCheck = false;
    }).catch(function(err) {
        console.log("Get axios excelDownAjax err : ", err)
        excelBtnClickCheck = false;
    });
}
