package config

import (
	"fmt"
	"testing"
)

func TestPacket(t *testing.T) {
	htmlCode := `
<?xml version="1.0" encoding="gbk"?>
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
    <title>rechargeNextday测试报告</title>
    <meta name="generator" content="HTMLTestReport 0.1.0"/>
    <meta http-equiv="Content-Type" content="text/html; charset=gbk"/>
    
<style type="text/css" media="screen">
body        { font-family: verdana, arial, helvetica, sans-serif; font-size: 80%; }
table       { font-size: 100%; }
pre         { }

/* -- heading ---------------------------------------------------------------------- */
h1 {
	font-size: 16pt;
	color: gray;
}
.heading {
    margin-top: 0ex;
    margin-bottom: 1ex;
}

.heading .attribute {
    margin-top: 1ex;
    margin-bottom: 0;
}

.heading .description {
    margin-top: 4ex;
    margin-bottom: 6ex;
}

/* -- css div popup ------------------------------------------------------------------------ */
a.popup_link {
}

a.popup_link:hover {
    color: red;
}

.popup_window {
    display: none;
    position: relative;
    left: 0px;
    top: 0px;
    /*border: solid #627173 1px; */
    padding: 10px;
    background-color: #E6E6D6;
    font-family: "Lucida Console", "Courier New", Courier, monospace;
    text-align: left;
    font-size: 8pt;
    width: 500px;
}

}
/* -- report ------------------------------------------------------------------------ */
#show_detail_line {
    margin-top: 3ex;
    margin-bottom: 1ex;
}
#result_table {
    width: 80%;
    border-collapse: collapse;
    border: 1px solid #777;
}
#header_row {
    font-weight: bold;
    color: white;
    background-color: #777;
}
#result_table td {
    border: 1px solid #777;
    padding: 2px;
}
#total_row  { font-weight: bold; }
.passClass  { background-color: #6c6; }
.failClass  { background-color: #c60; }
.errorClass { background-color: #c00; }
.passCase   { color: #6c6; }
.failCase   { color: #c60; font-weight: bold; }
.errorCase  { color: #c00; font-weight: bold; }
.hiddenRow  { display: none; }
.testcase   { margin-left: 2em; }


/* -- ending ---------------------------------------------------------------------- */
#ending {
}

</style>

</head>
<body>
<script language="javascript" type="text/javascript"><!--
output_list = Array();

/* level - 0:Summary; 1:Failed; 2:All */
function showCase(level) {
    trs = document.getElementsByTagName("tr");
    for (var i = 0; i < trs.length; i++) {
        tr = trs[i];
        id = tr.id;
        if (id.substr(0,2) == 'ft') {
            if (level < 1) {
                tr.className = 'hiddenRow';
            }
            else {
                tr.className = '';
            }
        }
        if (id.substr(0,2) == 'pt') {
            if (level > 1) {
                tr.className = '';
            }
            else {
                tr.className = 'hiddenRow';
            }
        }
    }
}


function showClassDetail(cid, count) {
    var id_list = Array(count);
    var toHide = 1;
    for (var i = 0; i < count; i++) {
        tid0 = 't' + cid.substr(1) + '.' + (i+1);
        tid = 'f' + tid0;
        tr = document.getElementById(tid);
        if (!tr) {
            tid = 'p' + tid0;
            tr = document.getElementById(tid);
            if(!tr){
                tid = 'E' + tid0;
                tr = document.getElementById(tid);
            }
        }
        id_list[i] = tid;
        if (tr.className) {
            toHide = 0;
        }
    }
    for (var i = 0; i < count; i++) {
        tid = id_list[i];
        if (toHide) {
            document.getElementById('div_'+tid).style.display = 'none'
            document.getElementById(tid).className = 'hiddenRow';
        }
        else {
            document.getElementById(tid).className = '';
        }
    }
}


function showTestDetail(div_id){
    var details_div = document.getElementById(div_id)
    var displayState = details_div.style.display
    // alert(displayState)
    if (displayState != 'block' ) {
        displayState = 'block'
        details_div.style.display = 'block'
    }
    else {
        details_div.style.display = 'none'
    }
}


function html_escape(s) {
    s = s.replace(/&/g,'&amp;');
    s = s.replace(/</g,'&lt;');
    s = s.replace(/>/g,'&gt;');
    return s;
}

/* obsoleted by detail in <div>
function showOutput(id, name) {
    var w = window.open("", //url
                    name,
                    "resizable,scrollbars,status,width=800,height=450");
    d = w.document;
    d.write("<pre>");
    d.write(html_escape(output_list[id]));
    d.write("\n");
    d.write("<a href='javascript:window.close()'>close</a>\n");
    d.write("</pre>\n");
    d.close();
}
*/
--></script>

<div class='heading'>
<h1>rechargeNextday测试报告</h1>
<p class='attribute'><strong>开始时间:</strong> 2017-08-01 14:13:37</p>
<p class='attribute'><strong>运行时长:</strong> 0:01:22.262288</p>
<p class='attribute'><strong>运行结果:</strong> Pass 2</p>

<p class='description'></p>
</div>



<p id='show_detail_line'>显示
<a href='javascript:showCase(0)'>概要</a>
<a href='javascript:showCase(1)'>失败</a>
<a href='javascript:showCase(2)'>全部</a>
</p>
<table id='result_table'>
<colgroup>
<col align='left' />
<col align='right' />
<col align='right' />
<col align='right' />
<col align='right' />
<col align='right' />
</colgroup>
<tr id='header_row'>
    <td>测试文件/测试用例</td>
    <td>Count</td>
    <td>Pass</td>
    <td>Fail</td>
    <td>Error</td>
    <td>View</td>
</tr>

<tr class='passClass'>
    <td>recharge_day_rescure.json</td>
    <td>2</td>
    <td>2</td>
    <td>0</td>
    <td>0</td>
    <td><a href="javascript:showClassDetail('c150156809957',2)">Detail</a></td>
</tr>

<tr id='pt150156809957.1' class='hiddenRow'>
    <td class='none'><div class='testcase'>case_01
对昨天点击退账，验证可以正常退账，验证cptran表和cpctran_intermediate表补账数据正确</div></td>
    <td colspan='5' align='center'>pass</td>
</tr>

<tr id='pt150156809957.2' class='hiddenRow'>
    <td class='none'><div class='testcase'>case_02
对不同账号，多条昨天点击进行退账，验证可以正常退账，验证cptran表和cpctran_intermediate表补账数据正确</div></td>
    <td colspan='5' align='center'>pass</td>
</tr>

<tr id='total_row'>
    <td>Total</td>
    <td>2</td>
    <td>2</td>
    <td>0</td>
    <td>0</td>
    <td>&nbsp;</td>
</tr>
</table>

<div id='ending'>&nbsp;</div>

</body>
</html>
`
	tmp_report := []byte(htmlCode)
	for i, val := range tmp_report {
		if val == '\t' {
			tmp_report[i] = ' '
		}
		if val == '\n' {
			tmp_report[i] = ' '
		}
		if val == '"' {
			tmp_report[i] = '\''
		}
	}
	htmlCode = string(tmp_report)
	str := fmt.Sprintf("{ \"service_id\":\"11125\", \"pms_id\":\"110\", \"duration\":\"60\" , \"interval\":\"10\" , \"master_slave\":\"master\", \"policy\":\"random\", \"schedule_status\":\"create\", \"task_result\":\"running\", \"error_step\":\"\", \"operator\":\"tester\", \"report\":\"%s\"}", htmlCode)
	//str := `{ "service_id":"11125", "pms_id":"110", "duration":"60" , "interval":"10" , "master_slave":"master", "policy":"random", "schedule_status":"create", "task_result":"running", "error_step":"", "operator":"tester", "report":"abc"}`
	//fmt.Println(str)

	cfg := NewParaConfigSingleton()
	ret, err := cfg.Parse(str)
	if err != nil {
		fmt.Println("ERR------------------", err.Error())
		//fmt.Println(htmlCode)
	} else {
		fmt.Println("RET------------------", ret)
	}
}
