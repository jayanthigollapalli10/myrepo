def xor(a, b):
    result = ""
    for i in range(len(b)):
        if a[i] == b[i]:
            result += '0'
        else:
            result += '1'
    return result


def divide(dividend, divisor):
    pick = len(divisor)
    tmp = dividend[0:pick]

    while pick < len(dividend):
        if tmp[0] == '1':
            tmp = xor(divisor, tmp) + dividend[pick]
        else:
            tmp = xor('0'*pick, tmp) + dividend[pick]
        tmp = tmp.lstrip('0')     # remove left-side zeros
        if tmp == "":
            tmp = "0"
        pick += 1

    # Final step after loop
    if tmp[0] == '1':
        tmp = xor(divisor, tmp)
    else:
        tmp = xor('0'*len(tmp), tmp)

    # Remainder must be (degree) bits
    return tmp.zfill(len(divisor)-1)


# CRC polynomials
keys = [
    '1100000001111',        # CRC-12
    '11000000000000101',    # CRC-16
    '10001000000100001'     # CRC-CCITT
]

print("Choose the CRC")
print("1. CRC - 12")
print("2. CRC - 16")
print("3. CRC - CCITT")

n = int(input("Enter your choice (1/2/3): "))

send = input("Enter the string of binary data bits to be sent from the sender: ")
rec = input("Enter the string of binary data received at the receiver side: ")

# Select polynomial
key = keys[n - 1]
length = len(key)

# Sender side calculation
send_extended = send + '0' * (length - 1)
remainder = divide(send_extended, key)

print("\nSender CRC Remainder:", remainder)

# Receiver side validation
ans = divide(rec, key)

if ans == '0' * (len(key) - 1):
    print("No error in transmission")
else:
    print("Frame error detected")


Choose the CRC
1. CRC - 12
2. CRC - 16
3. CRC - CCITT
Enter your choice (1/2/3): 1
Enter the string of binary data bits to be sent from the sender: 1011
Enter the string of binary data received at the receiver side: 1011111000001111


Sender CRC Remainder: 010001001101
Frame error detected




Case 2: CRC-16
Choose the CRC
1. CRC - 12
2. CRC - 16
3. CRC - CCITT
2
Enter the string of binary data bits to be sent from the sender: 101110111010101
Enter the received data: 1011101110101010100110011111011
no error







