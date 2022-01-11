# sap-api-integrations-employment-information-reads
sap-api-integrations-employment-information-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 従業員雇用情報 データを取得するマイクロサービスです。    
sap-api-integrations-employment-information-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-employment-information-reads は、オンプレミス版である（＝クラウド版ではない）SAP SuccessFactors API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/ECEmploymentInformation/overview   

## 動作環境  
sap-api-integrations-employment-information-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-employment-information-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-employment-information-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/ECEmploymentInformation/overview    
* APIサービス名(=baseURL): https://sandbox.api.sap.com/successfactors/

## 本レポジトリ に 含まれる API名
sap-api-integrations-employment-information-reads には、次の API をコールするためのリソースが含まれています。  

* EmpEmployment（従業員 - ヘッダ）※候補者の詳細データを取得するために、PersonNav、PersonalInfoNav、と合わせて取得されます。
* PersonNav（従業員 - 従業員情報 ※To）
* PersonalInfoNav（従業員 - 個人情報 ※To）

## API への 値入力条件 の 初期値
sap-api-integrations-employment-information-readsにおいて、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.EmploymentInformation.PersonIDExternal（従業員ID）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Role" が指定されています。    
  
```
	"api_schema": "EmpEmployment",
	"accepter": ["Header"],
	"person_id_external": "109031",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "EmpEmployment",
	"accepter": ["All"],
	"person_id_external": "109031",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetEmploymentInformation(personIDExternal string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(personIDExternal)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```
## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、従業員雇用情報 の ヘッダデータ が取得された結果の JSON の例です。  
以下の項目のうち、"PersonIDExternal" ～ "PersonNav" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-employment-infomation-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-employment-information-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"PersonIDExternal": "109031",
			"UserID": "109031",
			"LastModifiedDateTime": "/Date(1442368786000+0000)/",
			"ServiceDate": "/Date(858556800000)/",
			"EndDate": "",
			"BonusPayExpirationDate": "",
			"CreatedDateTime": "/Date(1442367865000+0000)/",
			"IsContingentWorker": false,
			"InitialOptionGrant": "",
			"InitialStockGrant": "",
			"EmployeeFirstEmployment": "",
			"EligibleForStock": false,
			"EligibleForSalContinuation": "",
			"OkToRehire": "",
			"AssignmentIDExternal": "109031",
			"ProfessionalServiceDate": "",
			"SalaryEndDate": "",
			"SeniorityDate": "/Date(858556800000)/",
			"StartDate": "/Date(858556800000)/",
			"HiringNotCompleted": false,
			"IsECRecord": true,
			"PrevEmployeeID": "",
			"RegretTermination": "",
			"CreatedOn": "/Date(1442353465000)/",
			"LastDateWorked": "",
			"FirstDateWorked": "/Date(858556800000)/",
			"OriginalStartDate": "/Date(858556800000)/",
			"PayrollEndDate": "",
			"BenefitsEligibilityStartDate": "/Date(858556800000)/",
			"StockEndDate": "",
			"BenefitsEndDate": "",
			"AssignmentClass": "ST",
			"LastModifiedBy": "admindlr",
			"LastModifiedOn": "/Date(1442354386000)/",
			"CreatedBy": "admindlr",
			"PersonNav": "https://sandbox.api.sap.com:443/successfactors/odata/v2/EmpEmployment(personIdExternal='109031',userId='109031')/personNav"
		}
	],
	"time": "2022-01-11T09:49:19.553754+09:00"
}
```
