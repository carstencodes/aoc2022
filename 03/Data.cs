namespace aoc2022.task03;

internal static class Data {
    internal const string Sample = @"vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";

    internal const string RealData = @"hDsDDttbhsmshNNWMNWGbTNqZq
VQfjnlFvnQFRdZWdVtqMGdWW
zvvvRnFFfjjlRBlBPzgQgRvvmtrmhHcptLHCDhcHHmLsBmsB
FrzFvvdTDcTnmTzdDTTzdvWmjhgVPrhSljSQSPwPjPjPjSVC
sMsGbqGsbbRqRbBMBGRMbLpNSSpjhlQljHVClhjgPjjPhlVp
sNbGtJbMfssNtvcnWFVmnvDd
TNfmdFJmfdZMQffVRQVV
jVHBCcDSjWrMZjvg
SShSbCGpcBtBtwtVLJJddmtLmT
CtpNftbNWbtSJDHqGZJFLfLr
dPsHlsRBHcZdqDFDZwwJ
snjVlvTPlPjVlQlHWjpSmzgNNzSmtpSm
qhZtSVqCqThGcGzZnnfZcB
WbddWbDwrBzcpzHpBb
DBBMFWRJDrDFWLWljCqjQjFvtCsqTjqs
vhFTzRzzTmPvbplWFtQttQQZtZhMZqcqSQ
fJVCfDfJNCLDwJNGmssZgwqgZcmtgcms
VmdVNLHGGVDBdfLCHnLGHnbWpTplWbddRTlzplWPpbFp
smwtNVqRjNmZjZBDSvzSzl
FnTJFcTTFccCrJGTLncdCCcPJZfBBDSlSJwZDlggSffvgSSf
FWTFGLFWLCPWrCnFnQWNbQVVwphhmHHbptVsss
BrrgrtgfBpPFhhgMWq
ZGvsvDGClvsSRScpGBhPphWMPhhTNh
SBSBdBZCcdwHrVQwQr
qmFqdVtqsVdzqGbwMJwGPpmPHM
ZjTjLQLLDrrLjcFhlfrGHppfbJwGpMMpGHwRCC
BTlZjrBBBcLTcDrjBlThWBjBtWNqnSzSnsSdvsFggNnqVVnv
zVvmjGgpcJnbTTTJHRHSRb
NPFrFQfCLPrdRlbtQRRvBtHb
frMLqPFMrfrPLCwqqNvjczwwnmwggGmssnmnnm
HbFJhhshsffcvslmGmLFrQBrlFTG
jNRPwwPjSPCdCdvRzRBTlRzmGrmB
nCSWdNjCCPqvtttZnDscth
ScSrRTPcSSDRWSptWcdmmWGbmGGLmLJvNJNbbJ
flzHjFpZjFfjjgszjlqzJNnnvsmbMmGvJLNNbmJv
jFZFVpffpVlqfQhgtTwcTVtrTcRBwBTD
rHrdGSMSSbZbjShj
qZfDBBvllvvWLtqbbQwhJjbtwbnQgN
WzWlzZmLWBLZLzCzrHMVcRrMRFRCCccF
BzdplppDlBBrqWnjFMBWqNWq
whZhZSSHhhVSrvSgHPvgvjnFTPsFFnnNTcjTFnTTsc
HLCwVSfZLffSHhLvwtQbJbrdGlRRrGdmpztG
grDFfDlfCftCzCfNztclNFrBNQjbZjJjjPPVsjNvsbvPsj
HwwTpGpRwMdpHWhvjzVsPJjJGVvVZs
ShphRpwzMWdpHdwWHwMpnHwLDmmgcLCCDtLDCCSlcgcfCD
ccqqLLqCqTSlZMLQMllZTvnNfjddttmmDpRJjvhfpfthRmdf
rbVssWwggFrGsWzPbVFGJpftQRPDmQDpdPmdJmfR
bBHWWHzWHWWHrsVbsFgwbFqqMLTnBclTMMSnLnqLqMQZ
CcSPGCCPrdPtdjcsBLDghbVLhqDl
vMJwTHzPvzVwqBBDblls
QNvMFfRMJHPZHjnfWmdftSjSnp
dnBCPhhBCrQfChdbNVGLszzDzVDsTbWT
HgcJgpppPqqHwPwJSczVzDssNNVWTtqVVGts
JHRFpjFccplcRwPJpHPScpMPQmCdBQQfQjjhCfBCrQQvdmnn
rQGmVRLRbDRHmmZLGBGVLHBVFspSstWWWNJcsgpQTSsNppJS
qlldhPdfCgnspJFWCFsFNT
jMMzwndfhnwPfqPMjgjMVBrGBmrHDHLZbjRGHbHj
wzpZfzHRSRfzgHfffZwwStCtSrBhBBCTrtFhhBFG
QPjQQQDcDWJNFWtrtWrGmTCMtmBW
lDvvDQcdjQcQLvlDwnpFgbbznZZglZsl
RfMFTMFrVrSRFPlFSfVlHpLqgzpHBLzHBBVzVpHG
CchbhcwdmdJmwJJtGgnqzppLmGGBQqQp
hCsstCJwLvMRvsZsTR
fQlfMlNClQhhZhrlWrWw
njDbnTDTBtGjmrGvSh
bgshBdBcDbTTdnnnTqcqLgqfpfQppCsCMsHHVVpHHNCHFF
PbCnTbzJnqQNzbbTNDdpwcmjDmwjGQjccw
hWgvSdLvwcGjSpSm
vVfrFvvhHFTZndJq
FFvRVCRqVRcfsDLrgqGNWjjHfhQQzGWjQHzN
ppJPBwplwSBJTmPpTzWWStzHHjNNLNzHNh
wnwMPbLJMJllJJwBmJmnLVvCvsCsbFgDrRrsCsvrqc
jqHgVgdgGQttWCtNqNflmllgFnfDnmFFlpcl
TZZsrrwwhwrsrZRGmhcfSnhGlmSFcf
rsLvvbJPPLBGPCQqHWtMCMHN
WzzBpCBpMsBpCvCfsgnPPfHgbfFNfF
jtdTLLjGTGjDjLbbbGlDLLLmfFgmmfgrPrgmNPHSnnNnqFSq
jbRRbjlwlRVpvWwvzBpB
qpwzCzCznFznTcCvrcrvVcLb
cPmNMHSlMsLfvWgsrWvL
mMHGPDMBGSGPHlPBcBPzznnzpdQFBjqRFdwdnn
QGZLJzmJrZgZzZhNQFqDWlWPWDFCWNRlPW
hMMbhVbhHWCsPWCMRs
BwjbSHVBVvfcTfZgZzQhrzdGrt
cvPTjfDPpDmmBjbQjZMdlBZj
CHNnghNChVzNgrFVwCMJLMMMMMQQdbLZ
FNSzShrHhNnWgVnWfvfbpfpDGTfvsG
FFpVrZhpTlSQlQzTtRtZHfmPmJDbRtZJ
jNnwBLnWwBgNBNCwsNgsMsCLVfDWfmDJRffmRRtmfRmPDRtR
BnLdNdjnBNcLngdCVBndgllFlQrrpQlSSFrQphFcQq
JmVLJPMNjmVJpMLJSVmNQZZQZZrnTHqZQHTrTTMr
ltfdwChhwRdRswDDdnBQqqWNTqrrHrqZRr
dwdwwDwsGlhhDFtsCwhsJLcPmGPcGzSjSjjLzNPS
QgSbgQCLQSJFMccLFLVVzH
WBNffrBpBNdNRdWDfptBtdzWcMZZVPMVwHMmsVHFccHsDsVc
ptphRWrfGRGRnqlCSQvhqbCzJS
VvdMLMLMBMlVlVschsNpDGpdNsGc
tqFSmnmnnttGfDqNcfvNDD
nzRHnrwrrWRrrzHtbMlTMBjCvWLgBBMl
pCBlRvzwzlCzvZqqDwzmvgtsLsQdgZsgPNtdrsWrst
JbGjbGVGHSFbhbnhTShSTbQtQrPsLsHLQgNRtsgdQsLP
GnnbJbGMGbSjjbSbCzqlwMRRDzqzMBBv
TTVRJVMWMshSQtjSVTQJRQlcCBncJccdppnJcBBDngFpgP
frvfzfHrwzZNrtNwzzzZncCCZFdFFCgBDpPcBcBg
LwNGbqwvmvvtbNQjQlVVRWTGTQWM
RnggwVLRLDfCVZhfpDGGMGMGcGzGNHvv
jmmWBTSsBmFmSzctsqpccHvzpN
SBSrTblSQPbQhNwwZfPVdZPf
CCvCwzfNStLzfrbmMJbZMtlsbJMW
gPPPBqDjBcPFpVgBRbnMsVsbJZnWsdbSbM
qgDPHjHcPhpDRRpPBRpCGSLwvHQwGfzrHLGLTL
CLGqDZZLTdddPsdJpq
gbRbbnghnrWvgrdJdSTRSsVNJlld
hMnwrjnnjggvnLDwGffTfwCZZZ
NzJHbNHNNzJzgmHmzpQSvvLqbLsVVsVGvB
WtWhtWDdrZldDWrWTlZgppVVsqQTVQBqsGqBsQVp
jWjWRRRlPcHRwJgw
CCnnFTmnPCMCRNfnwGwdfzvwwl
VQQVShDSSshhDDtDLhjccGjLBBzBzlZflNZvwZzdwBzpSNNZ
QLHhJDDhDhgscgtjbGHTrWbrTbRmbFrm
CJbLvJvbwtFHqvLzwJqqqtHWTWRgDScDRSWQQjTRcWRDLT
mssGsMNphZMNsPPBnhSjRgdnQRdWgdjgrcDn
GGmBMMsffmslMGshZlMphGqCzHCbzlbvzqwzzgFJggCF
WCgWBphpWLQZQpgdhGdwmfbfFRVRjRTbbSFttdbSbT
qqrZnDNqZJDTVzRjVFbfSN
rPMvqJJqrJMPJnZMZZgLPgLWLggWQghwhmBh
CWGGzdHHmPPSmPsC
LqwlZwRLrPMQlMqrlbZrQRsSNsmssSNSsNcBNpgmsJ
lwLDQhrDMQqPGfhzGGjhGn
ZqDlZssCqJJMvpdBpBBmBQSMRp
wLgVcbgFLzTLTNNZmNNRdjdRmF
HZHbctWTwgVWgsfrnnPqlnsWlD
RSnwSPFcLnFPnRwjzctzbGNlZgNbbGdGpLhZdpgM
BqqBfMvTmmJqDgGNVdGVbVJJZG
vsTDfqBmHmMWQCwjtrHjjSFFnRSn
LsCmmcDHRjdtNMstwwzJ
TvThqfBFBNTnnndTtL
lvGQfbFQGblFrRccLRSSlPPVHS
qbLpqTHSqpbqbrPcQgjPDjcdDL
gnzhhBBwBWZzMglmjDrDPjvfdvQPdwtd
WZZzZZlZmhsMmFgRBBBzHHJpNJGVRSJVGTVpNqCH
JDphhGhDdGzWRBnvqqLDNLMnLw
gsrTHHffTHPcrPrlHCNZhvBZnZNLhPvBNwvh
sHVCSsSghJpSJjQh
JTMGlfjlTdqjnqbnqFwqmnbQ
PBZhBBcWRZprPZcZDDCZTZRgnnzwbbsbnvhbvznFsNFFvvNQ
BcWrVgCCZTDRDrlGttfVtMddSJlH
vwwvpVbSvnSRRmfMCmTHVHTBHB
QLZgDPgSDgGTMZfmTTBZ
QDDsQFDlzlgtJlLdFDgSJQFvvpRvjqzjRwwwzWvhvWwqjj
mRRTGGNNflGRGGmmgRblsGwCZwVZlZjVwjztpjZhpBCB
PMLLFLHPLPnLqDDLvFDrzzMjhwVCjtphBzjMhV
FDdSPSpLcDsNRRWSTNWN
STldJthdJbtTqljCRDDHmqmj
VVvNwwvNFssJFJPNNwVvRMCgCgDqjjjqrDqqMHqP
QBZwQfZwfVhtcSBtBJnT
TzjjPzsQTslNlNzPRVGJJJGGtTJmgJHtmTZC
dBDWScMBhhPGgdwwJPfw
SqSqbSPDBhqnMqvrrSWVNFpRVRLzVQslvpNjVL
bWFgFCPFtgvDZWgtChDNFJHvGVzHHpjzHnnzGzzHRR
qcScQbbmqdQmlQmrlcQwLmHlRRjzGHnHJnnGVjHzBHzG
TqQwmLmfcddfwrfCgbWCNPsZNfCb
pddprrtrCPdvJdMjwwwHnLwwjLWCLg
qhzZTmZcmRhmpFlVHcQQVwWQHVQwnH
lGmhfRfmBZRlmmbvDPBMvNbvJJpP
NsptgfGLLNwnNQSZbCvZnRnMCb
JldhdzwzBMCSZvrz
JFcdWTdwhPTFVDVmTJNqmstLqgLtLtjGGpsG
dVVTSgTDpHVDjgdWpdpHTZSbWGrnnvrNwzFGNrFwnNNwvh
CPRlMPJcMQcBcsmmLCMPrzbFfhwfrvLrNNwwGwfF
CRJmtbmJlQbsQlRBpZDVjTHTdjDtSjZt
rQVJrRFdrwDfzHQHQBTnpWTW
PCLbPcPCsgqCgPgLjScSqNbHTzMtWmWtzlTHmBtTlMssMT
cCqghSSPcvgScPbwFGdDDVZFfDhZGB
zrRQRdqzPHQtnMPrtzPMRRQMVBBblJJBSClBpJbpdCCbBlCC
hTcGwzswGwGmGfDvvfGmGNfBpllVSWbWppNCBNBVpCBClW
gvGFTmTgwDhTDccsTfzfmfGGQPgPPqrgRZPHnRqRZrQLnRgP
hvmmJllPbmCRMNGMMlNwNl
PFTpTVjTgpTpBRgMGMnRNHBB
WWrqzTTPVQDPqpjTqPJbmLtcfsQsftbLbvct
SzrmpjjcsjTZNzgnnNzN
BLHNDwBLBPLwLBhwDVLgdQCgCQGTngHQZCngZd
PPJBDvBVVBmppNjJjrrr
ZHBNQFhsqHBsgCfqtctcPvSwPqrV
LlnGTnJpJJTmdDpmLlmLndWfVrPvvRwDfcwwwRVwcfQtvP
GnbQblWmWGdTJQdTGnZHsHhZhFNsbCsjFgjC
hWfDzDTVndDMhddMlBWMBDfJRnRtvvSSQjCvZCtjtpJvSR
bGHsccFcbscsqGPHNGcrpjJZtvSRtFtQCZrjSj
GsbwGGwNNGLgPLwMzBzfMMVMTLdTCC
GBcNzTSSmGzmTLNgvwgpNCDqpDggpw
JRZMrJWFZZnZtJgvvjwbpbCJDd
rFMPRhZtZFnWrRtQGmPPDcLfmGLTfz
VdWnVdjhhdFjVWbndMlNLQspVMHCNVlClV
RSrJBRRJwJSBQpMBHLLDCL
TqwtRRRJzJTSqJSzSrtmqgWWhcncvPgnWbPQnbnWmb
VnDFpPpFssVSpFDVHbRbscCvgbMTvTCR
JfzqdQBfhBdddfBBGDLdGQvbrqMMcCRRMTgbqgMrbbqc
QDfzJNWBJLQBhmdGDzDGhQGGlFZwPtWjtFFppllSVpZZFnjj
qrLLNpJbJnRLNnpvQtRVhhRFCdlFFlFd
mmjzjvGjwPwmTsSTSQjDVlVWQjlCDthCCC
cSSmcTTPcSswScfSHmTSTzJqqNrnpBpqBbJLvZMrqfrL
NSvRZRfFvfHSZQcNJBLbzDLnrDFnhtFLFnrh
wmTGpmGCwsMplMsHllPlMnDLjznrgrzDjgnntznr
dsCVGGGwmpTGPplmCmPppVmHSSRJNfJvBNZQfWdBRJZRBZcR
TwQwqDPQtwNwzNDTZcnZbJvMnMMbFqZM
SzGSjrjLWrjHHspWVhvVVnFJbccVZcRJbllb
pHppszGSprhhWHLCLrsjdTtDDPfwdfwtdNfDgNCN
ftcvBtBFtmBlmvPFmmcczCChrgSCzzCSnCSSnGHf
sJddbdTDbDHdnJRggrGzGzrG
dppDVDZMMMsTTVsDTsTDpwVctNcvBZQPcPctqtQcHmvlvQ
jzbdzztbDqNqwvLvRmQZjvRH
FSJbFFWgJnZFLRZmHmRQ
TgVJTVSJGJcJlllgTMdqpdNsrztNNsNbMDDp
CCCVWbwVnlRbTcqSShqGhhGcnF
PgDBfDpMNlfgpPfNZZtcJgcqqhmmjqSmjFmhmS
tpfpsPrlpsPDDMDfBZrwLrVWLLLWRCdHLTwbVR
pjvfDGjSMpvDmDpDpSDnJmfqbPVsCMFsPqFVPqCrwrbMFV
NQlHtHNhZHgZZNBHhQgzPmCwbqqVFlsrPFrFCP
hgHtQdQchcHctHgcgNgBQdWNpmvTWvpGmLJDLGjTpLGnnjfv
QhgLLLmtlRqDtRGP
HLbnCZFWVHLZnFCJJRFrGJzDGDGJDD
WZHfndfMfCZbMnTVTfZhSNQQpdwSdLwhNcmdSN
sPwrPMgLFPFFsLZtmcclSSZDtcZs
qVzqdNdCnnNVVNCGmbncDBlmBlBBnRlZ
VTdCGVvVfffrjpfMQPwm
BPDldDTDPZcggjcccTdNMbbMNSQNqqjtzMbrRb
LvmWsfvssLGnQbQMRQqrSRnz
WpvsVmmpmmfpfJGrHfVCHVvmcDgpDlZphgFgdhclhdgdBlgF
VGwHbNzMMrzHbbHChhqgCqPNghgCqW
ZJVBvBvZWqvRvggP
JBJlBlBZcsBfcJVrHnLwQQGzLQMc
gBWfBPPPfhvVWFfSVfVdjjbvTvwwQppHcHcctTcQTHcZ
DnNnMJMqMJzqchbZtTQQrb
llRmNLDLDGlCsWSFCffWdshd
LpNMZZpqqpfTTwNqLZwGsZqZbdHRHbHGddnCBHRcmzGmmCdG
JFRtRlVStjPlhtjbBzBncmVWdzWBnb
rPhhSlrvQlFFFPgtJlJtFlhlDNTwRMfZTZfDZNrspZLMMsrq
zBLjLFBjLjmHWlzNZlzVCC
dcJrdfddbllJbdMTwDNMZWNVwVDwHT
gRcgJbcbqfgbftdjlqLhFFLPPhGBjm
WfBgBRzQGNNQqmmqZN
nFjCjCpLbtpPJtCDDnCDJpzncrSVbmdVqbhhdqNbSSmrdVSq
CLPJpDLlLlFDpFjjsGRsBGRfWwsHHglz
lSlSlpCRSsWTRLTlWRvlmMrBPjBPjpqrrmqPJMPZ
DDzbhVhQhDGzhQnGGfnHHQGBPZjMqJjBJMBVJmqMdrqqdT
NNGQbFwnHzNzwbQwFnwbfsLCLtsvLsWggFslsTggSc
nvzPvCnlvtwCrZWmWwvvZCQfbbfQfGbqSJJGmqGSFSbJ
LhTBWdsMNNRgNcgDWsDNcVSfQqJGFSFJqSSddQGSJF
HNgchHcWDRNhTNMWwtPrtZZjnHzrnvCz
djhnzRghMMVCBfhh
qjQTrTPQJCDDqBDJ
LQvGrLjTHLjNNPPTpQgtztSmmbFgmgLbFnmL
FRDNFBBRRVFFmbLZHPZBZvvH
QnhgMllglJTdGgJnhLQQJpZpvwZHpwsPTwpbsZHmsH
lnhnQGrMgthMlntlGfQhgWWcRSDcVCrLWzRSrRFDRN
PqrrrRnPBbrVhVqFrFVRPVhZLvNSNvLZcQvtJfRvNScJNJ
dDzWwwCTmmdwdddpDLWQZMSSMfSJtcWJfQSQZN
CCwmTdjsClVjFjnLBl
srjCvjPmQVlPjFPmQmPrdHHZhvHZDqHhDDwHHqfB
pLcnJQNQMZpqZDDZ
WNRbtNJgRPjjQVmz
NJJRmjmJbbJfqSVMNHFCSFzLLlrLLrFHTz
QvnsQGvBwWwQvgRHlGGDFPFCGlrR
QhvwBvBctBccZWZNRNmVfjpmjJjb
RMmGGMLRRCFmRPPfGFpGPFPJWZQWctrtlQvZvltfrQWcWWBq
gggwjjbjwwbZtwZBBcmQQv
SdNbDDVSgPMFmPzdMm
nZhnNZDnZPmZPWbppPpMlvRlzvrtMmRtqRzRfq
HcFwsCQLVQwFwLtLbvtzrlrLtt
GsgCFCgCQHHCVHsFQHcFdDPDbJDZTpZDbWJPNWWZDd
BBrBrGlGpgGjsNhlBlpBwpfSwZJdQwfcZwvSQnnn
LvWvHLmmVJQQHfQH
RPLRMvqFTbRTjGBhjNFsslls
cNZZZmZDcDDJmhzzrrlHtSbvgjSvgfPSWvPfjShv
VBwnndnVCqbqpRRpnspnqRWtGgWSSgvFBSGGSWgtGGSP
LqCMnTLVRwCRCpRLpbHDNzMMNcmmHNHQJQ
MMqDtnVnBlHtZvtB
WLWrWgdWwdrLCTFCwLlbbsJsJQsbQlQzlvrB
jFSvTdjfnfRmVcRR
ZLGqnvnqLzvbGRMfcRpwMpdV
fgfNNfgHHjVmRcVdgM
HsWDCDfCQCZBBZnvWtLq
bTZjqflqZhcrlczGzppGNgjmFNnp
PmmRSWWDMBQVNpWFznGF
SStRBDSCCSSSwPBwBDBwPmZhZlfZhqHTsTfltHHZfsHH
GbNbsSptQGqsdJCzsddcgzzv
DHRRnmWWmZnmRhllnHnnnMLvvLgcTVvjVhCTvgzcJgLj
RnWMlDZRlnHlmHWBFwGQqNGGPNQzPGqFwz
vSGvHpJnBLbGHBNCgfDzzChDgbCfzT
wFRslqmqTRgggQghPmQf
qjRFMjWqNNMMGHTL
fWGcQGGSRFQZhttZJfSSJflDDrwdClljVrNDdrdCFBCr
MTgvLLPPnHzMbDwdlNbMBwMM
mnTvnnPTcNmmJJWN
qqbbQQnbWrqGgnWqvZpVzMCZjCgfjZCSVM
ldcmDPDhmlFBHPDddLBVFDHLppZpjSCjjNfwNMwCpSMwhCMp
FtDdsHPcHmdHVPLtHsdtBHQnsbvnTRRTRsRRqbqvqWnJ
hhtBtPrgbbhhgjZjjCCHHNpNDHpffHWCvr
LGFLVwswsJMSgFwMMpddSvpHCCdDdvCpvm
sGsFsQLsVsLFnnFTJQthjcjQqhRcBZZtRg";

}