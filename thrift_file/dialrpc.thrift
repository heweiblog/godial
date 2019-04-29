namespace c_glib dialrpc
namespace java rpc.yamutech.com
namespace cpp rpc.yamutech.com
namespace * rpc.yamutech.com

exception Xception {
  1: i32 errorCode,
  2: string message
}

typedef string ObjectId 

enum RetCode
{
  FAIL = 0,
  OK = 1  
}

struct IpAddr
{
  1: i32 version,
  2: string addr  
}

enum TaskEvent
{
  IDLE =0,
  RUNNING = 1,
  SUSPEND = 2,
  FINISHED =3,
  CANCELED =4
}

struct TaskProcessArgs
{
  1: TaskEvent event,
  2: string batchno,
  3: double percent,
  4: double dialLocalRate,
  5: i32 dialAvgDelay,
  6: double detectLocalRate,
  7: double detectAvailRate,
  8: i32 detectAvgDelay,
  9: i32 totalAvgDelay,
  10: bool closed
}

enum DomainType
{
	A = 1,
	NS = 2,
	CNAME = 5,
	SOA = 6,
	PTR = 12,
	HINFO = 13,
	MX = 15,
	TXT = 16,
	AAAA = 28,
	SRV = 33,
	A6 = 38,
	ANY = 255	
}

struct DomainRecord
{
  1: string dname, 
  2: DomainType dtype
} 

enum DialMethod
{
  Dig = 0,
  DigAndPing = 1,
  DigAndHttp = 2,
  DigAndWeb = 3,
  DigAndVideo = 4,
  FocusDomain = 5,
  RefreshCache = 6,
  DomainSchedul = 7
}

struct VideoResult
{
  1: string  url,
  2: bool    available,
  3: i32  speed
}

struct IpResult
{
  1: IpAddr  ip,
  2: bool    local,
  3: i32  delay,
  4: bool    available,
  5: list<VideoResult> videoResults,
  6: i32 downloadspeed
}

enum FocusDomainResultStatus
{
  noerror = 0,
  formerr = 1,
  servfail = 2,
  nxdomain = 3,
  notimpl = 4,  
  refused = 5,
  others = 6
}

struct FocusDomainResultItem
{
  1: i32 priority,
  2: string value
}

struct FocusDomainResult
{
  1: FocusDomainResultStatus status,
  2: i32 delay,
  3: list<FocusDomainResultItem> results
}

struct DomainResult
{
  1: string     dname,
  2: DomainType dtype,
  3: bool       available,
  4: list<IpResult> results,
  5: bool       local,
  6: i32     delay,
  7: FocusDomainResult   fdr
}


struct IpSec
{  
  1: i32 version,
  2: IpAddr ip,
  3: i32 mask,
  4: string carrier,
  5: bool local
}

struct DomainTarget
{
  1: string targetid,
  2: string taskid,
  3: string batchno,  
  4: bool       available,  
  5: bool       local,
  6: i32     delay,
  7: list<IpResult> results,
  8: i32     avgIpDelay,
  9: i32     avgVideoSpeed,
  10: i32     totalDelay,
  11: i64   updated
}

struct AnalysisResult
{
  1: string     dname,
  2: DomainType dtype,   
  3: DomainTarget  home,
  4: DomainTarget  suggest,
  5: list<DomainTarget> totals
}


service Agent
{
  RetCode         registerModule(1: i32 moduleId,2:IpAddr ip,3:i32 port) throws(1: Xception ex),
  RetCode         unRegisterModule(1: i32 moduleId) throws(1: Xception ex),
  RetCode         heartBeat(1: i32 moduleId) throws(1: Xception ex)
}

service Notice
{  
  RetCode		  setConfig(1:bool enable,2:string host,3:i32 port,4:bool ssl,5:string username,6:string password) throws(1: Xception ex),
  RetCode		  sendMaile(1:string to,2:string title,3:string content) throws(1: Xception ex)
}

service Dispatch
{
  RetCode		  addDispatchTask(1:string taskId,2:i32 interval,3:i32 policy) throws(1: Xception ex),
  RetCode		  updateDispatchTask(1:string taskId,2:i32 interval,3:i32 policy) throws(1: Xception ex),
  RetCode		  removeDispatchTask(1:string taskId) throws(1: Xception ex)
}

service Analysis
{
  DomainTarget	  getDomainTarget(1:string groupId,2:string dname,3:string dtype,4:string targetId) throws(1: Xception ex),
  list<AnalysisResult>	getAnalysisResultPageList(1:string groupId,2:i32 skip,3:i32 limit) throws(1: Xception ex),
  RetCode		  reportTaskProcess(1:i32 moduleId,2: string taskId,3: TaskProcessArgs arg) throws(1: Xception ex),
  RetCode		  reportResult(1:i32 moduleId,2:string taskId,3:string batchno,4:list<DomainResult> resultList) throws(1: Xception ex)
}

service Collect
{
  RetCode         registerModule(1: i32 moduleId,2:IpAddr ip,3:i32 port) throws(1: Xception ex),
  RetCode         unRegisterModule(1: i32 moduleId) throws(1: Xception ex),
  RetCode         heartBeat(1: i32 moduleId) throws(1: Xception ex),
  RetCode		  reportTaskProcess(1:i32 moduleId,2: string taskId,3:TaskProcessArgs arg) throws(1: Xception ex),
  RetCode		  reportResult(1:i32 moduleId,2:string taskId,3:string batchno,4:list<DomainResult> resultList) throws(1: Xception ex)
}

service Dial
{
  RetCode         heartBeat()throws(1: Xception ex),
  RetCode         resetModule()throws(1: Xception ex),
  RetCode		  addIpSec(1:list<IpSec> ipSecList) throws(1: Xception ex),
  RetCode		  removeIpSec(1:list<IpSec> ipSecList) throws(1: Xception ex),
  RetCode		  clearIpSec() throws(1: Xception ex),
  RetCode		  addDialDomain(1:string groupId,2:list<DomainRecord> DomainList) throws(1: Xception ex),
  RetCode		  removeDialDomain(1:string groupId,2:list<DomainRecord> DomainList) throws(1: Xception ex),
  RetCode		  clearDialDomain(1:string groupId) throws(1: Xception ex),
  RetCode		  addDialTask(1:string taskId,2:DialMethod method,3:list<IpAddr> targetList,4:IpAddr sourceip,5:i32 interval,6:string domainGroupId) throws(1: Xception ex),
  RetCode		  removeDialTask(1:string taskId) throws(1: Xception ex)
}
