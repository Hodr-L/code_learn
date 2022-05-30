long_text = """
Variopartner SICAV
529900LPCSV88817QH61
1. TARENO GLOBAL WATER SOLUTIONS FUND
LU2001709034
LU2057889995
LU2001709547
2. TARENO FIXED INCOME FUND
LU1299722972
3. TARENO GLOBAL EQUITY FUND
LU1299721909
LU1299722113
LU1299722030
4. MIV GLOBAL MEDTECH FUND
LU0329630999
LU0329630130
"""
#设置分割符
A = long_text.split('\n')
#存储name和lei
name = A[1]
lei = A[2]
#检索title数目,以.开始以FUND结束
a = long_text.count('.')
a1 = long_text.find('.')
a2 = long_text.find('FUND')
a5 = a1
a6 = a2

# LU数组
c = long_text.find('LU')
c1 = long_text.count('LU')

# 定义数组存放内容
title = []
LU = []
isin = []
sub_fund = []
sub = []
SUB = []
fu = []
# title数组
#遍历title元素，并写入到title数组
for i in range(0, a):
    title.append(i)
    title[i] = long_text[a5 + 1:a6]
    a5 = long_text.find('.', a5 + 1)
    a6 = long_text.find('FUND', a6 + 1)

# lu数组
for i in range(0, a):
    "从fund开始到下一个.结束，将这里的LU放入素组"
    a3 = a1
    a1 = long_text.find('.', a3 + 1)
    # 从第一个“.”开始遍历，找到下一个点
    a4 = long_text.find('FUND', a3, a1)
    # 从第一个“.”开始遍历，到下一个“.”结束，找到fund的位置
    c2 = long_text.count('LU', a4, a1)
    # 从fund开始遍历，到下一个点结束，统计数字LU+数字的量
    c = long_text.find('LU', a4, a1)
    # 从从fund开始遍历，到下一个点结束，找到第一个LU+数字的位置
    for j in range(0, c2):
        c5 = c
        LU.append(j)
        LU[j] = long_text[c5:c + 12]
        c = long_text.find('LU', c5 + 1, a1)

    isin.append(i)
    isin.append(LU)
    LU = []

for i in range(0, len(isin)):
    if type(isin[i]) == int:
        sub.append(i)
    else:
        continue
str1 = "title"
str2 = "isin"
str3 = "："
str4 = '\n'
for j in range(0, a):
    sub_fund.append(j)
    e = sub[j]
    SUB.append(j)
    SUB[j] = ["title：", title[j], '\n', "isin:", isin[e + 1]]
    sub_fund.append({str1: title[j], str2: isin[e + 1]})

for l in range(0, len(sub_fund)):
    if type(sub_fund[l]) == int:
        fu.append(l)
    else:
        continue
sub_fund = [sub_fund[i] for i in range(len(sub_fund)) if (i not in fu)]

endfiles = {
    'name': name,
    'lei': lei,
    'sub_fund': sub_fund
}
print(endfiles)
