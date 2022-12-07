/// @use-src 0:"state/State.sol"
object "State_209" {
    code {
        {
            /// @src 0:58:2853  "contract State {..."
            let _1 := memoryguard(0x80)
            mstore(64, _1)
            if callvalue() { revert(0, 0) }
            sstore(/** @src 0:527:542  "block.timestamp" */ 0x01, timestamp())
            /// @src 0:58:2853  "contract State {..."
            let _2 := datasize("State_209_deployed")
            codecopy(_1, dataoffset("State_209_deployed"), _2)
            return(_1, _2)
        }
    }
    /// @use-src 0:"state/State.sol"
    object "State_209_deployed" {
        code {
            {
                /// @src 0:58:2853  "contract State {..."
                let _1 := memoryguard(0x80)
                let _2 := 64
                mstore(_2, _1)
                let _3 := 4
                if iszero(lt(calldatasize(), _3))
                {
                    let _4 := 0
                    switch shr(224, calldataload(_4))
                    case 0x07f2cc01 {
                        if callvalue() { revert(_4, _4) }
                        if slt(add(calldatasize(), not(3)), _4) { revert(_4, _4) }
                        let length := sload(_4)
                        if gt(length, 0xffffffffffffffff)
                        {
                            mstore(_4, shl(224, 0x4e487b71))
                            mstore(_3, 0x41)
                            revert(_4, 0x24)
                        }
                        let _5 := 0x20
                        finalize_allocation(_1, add(shl(5, length), _5))
                        mstore(_1, length)
                        let mpos := _1
                        mpos := add(_1, _5)
                        let mpos_1 := mpos
                        mstore(_4, _4)
                        let spos := 18569430475105882587588266137607568536673111973893317399460219858819262702947
                        let i := _4
                        for { } lt(i, length) { i := add(i, 1) }
                        {
                            mstore(mpos, copy_array_from_storage_to_memory_string(spos))
                            mpos := add(mpos, _5)
                            spos := add(spos, 1)
                        }
                        let memPos := mload(_2)
                        let tail := add(memPos, _5)
                        mstore(memPos, _5)
                        let pos := tail
                        let length_1 := mload(_1)
                        mstore(tail, length_1)
                        pos := add(memPos, _2)
                        let tail_1 := add(add(memPos, shl(5, length_1)), _2)
                        let srcPtr := mpos_1
                        let i_1 := _4
                        for { } lt(i_1, length_1) { i_1 := add(i_1, 1) }
                        {
                            mstore(pos, add(sub(tail_1, memPos), not(63)))
                            tail_1 := abi_encode_string(mload(srcPtr), tail_1)
                            srcPtr := add(srcPtr, _5)
                            pos := add(pos, _5)
                        }
                        return(memPos, sub(tail_1, memPos))
                    }
                    case 0x1a64b8e8 {
                        if callvalue() { revert(_4, _4) }
                        if slt(add(calldatasize(), not(3)), 32) { revert(_4, _4) }
                        let value := calldataload(_3)
                        /// @src 0:427:450  "string[] public dhtList"
                        if iszero(lt(value, /** @src 0:58:2853  "contract State {..." */ sload(_4)))
                        /// @src 0:427:450  "string[] public dhtList"
                        {
                            revert(/** @src 0:58:2853  "contract State {..." */ _4, _4)
                        }
                        mstore(_4, _4)
                        let value_1 := copy_array_from_storage_to_memory_string(add(18569430475105882587588266137607568536673111973893317399460219858819262702947, value))
                        let memPos_1 := mload(_2)
                        mstore(memPos_1, 32)
                        return(memPos_1, sub(abi_encode_string(value_1, add(memPos_1, 32)), memPos_1))
                    }
                    case 0x4985e85c {
                        if callvalue() { revert(_4, _4) }
                        let _6 := abi_decode_string(calldatasize())
                        let pos_1 := mload(_2)
                        let length_2 := mload(_6)
                        copy_memory_to_memory_with_cleanup(add(_6, 0x20), pos_1, length_2)
                        let end := add(pos_1, length_2)
                        mstore(end, _3)
                        let converted := copy_array_from_storage_to_memory_string(keccak256(pos_1, add(sub(end, pos_1), 0x20)))
                        let memPos_2 := mload(_2)
                        mstore(memPos_2, 0x20)
                        return(memPos_2, sub(abi_encode_string(converted, add(memPos_2, 0x20)), memPos_2))
                    }
                    case 0x4cd08d03 {
                        if slt(add(calldatasize(), not(3)), 96) { revert(_4, _4) }
                        let offset := calldataload(_3)
                        let _7 := 0xffffffffffffffff
                        if gt(offset, _7) { revert(_4, _4) }
                        let value0, value1 := abi_decode_string_calldata(add(_3, offset), calldatasize())
                        let offset_1 := calldataload(36)
                        if gt(offset_1, _7) { revert(_4, _4) }
                        let value2, value3 := abi_decode_string_calldata(add(_3, offset_1), calldatasize())
                        let offset_2 := calldataload(68)
                        if gt(offset_2, _7) { revert(_4, _4) }
                        let value4, value5 := abi_decode_string_calldata(add(_3, offset_2), calldatasize())
                        /// @src 0:1022:1120  "if (!verification()) {..."
                        if /** @src 0:1026:1041  "!verification()" */ iszero(/** @src 0:1027:1041  "verification()" */ fun_verification())
                        /// @src 0:1022:1120  "if (!verification()) {..."
                        {
                            /// @src 0:58:2853  "contract State {..."
                            let memPtr := mload(_2)
                            finalize_allocation_6438(memPtr)
                            mstore(memPtr, /** @src 0:1072:1082  "msg.sender" */ caller())
                            /// @src 0:58:2853  "contract State {..."
                            mstore(/** @src 0:1067:1108  "User(msg.sender, identity, peer_id, addr)" */ add(memPtr, /** @src 0:58:2853  "contract State {..." */ 32), abi_decode_available_length_string(/** @src 0:1067:1108  "User(msg.sender, identity, peer_id, addr)" */ value0, value1, /** @src 0:58:2853  "contract State {..." */ calldatasize()))
                            mstore(/** @src 0:1067:1108  "User(msg.sender, identity, peer_id, addr)" */ add(memPtr, /** @src 0:58:2853  "contract State {..." */ _2), abi_decode_available_length_string(/** @src 0:1067:1108  "User(msg.sender, identity, peer_id, addr)" */ value2, value3, /** @src 0:58:2853  "contract State {..." */ calldatasize()))
                            mstore(/** @src 0:1067:1108  "User(msg.sender, identity, peer_id, addr)" */ add(memPtr, /** @src 0:58:2853  "contract State {..." */ 96), abi_decode_available_length_string(/** @src 0:1067:1108  "User(msg.sender, identity, peer_id, addr)" */ value4, value5, /** @src 0:58:2853  "contract State {..." */ calldatasize()))
                            let oldLen := sload(/** @src 0:1057:1061  "user" */ 0x02)
                            /// @src 0:58:2853  "contract State {..."
                            if iszero(lt(oldLen, 18446744073709551616))
                            {
                                mstore(_4, shl(224, 0x4e487b71))
                                mstore(_3, 0x41)
                                revert(_4, 36)
                            }
                            let _8 := add(oldLen, 1)
                            sstore(/** @src 0:1057:1061  "user" */ 0x02, /** @src 0:58:2853  "contract State {..." */ _8)
                            if iszero(lt(oldLen, _8))
                            {
                                mstore(_4, shl(224, 0x4e487b71))
                                mstore(_3, 0x32)
                                revert(_4, 36)
                            }
                            mstore(_4, /** @src 0:1057:1061  "user" */ 0x02)
                            /// @src 0:58:2853  "contract State {..."
                            copy_struct_to_storage_from_struct_User_to_struct_User(add(29102676481673041902632991033461445430619272659676223336789171408008386403022, shl(/** @src 0:1057:1061  "user" */ 0x02, /** @src 0:58:2853  "contract State {..." */ oldLen)), memPtr)
                        }
                        let pos_2 := mload(_2)
                        calldatacopy(pos_2, value2, value3)
                        let _9 := add(pos_2, value3)
                        mstore(_9, /** @src 0:1130:1145  "identityMapping" */ 0x03)
                        /// @src 0:58:2853  "contract State {..."
                        copy_byte_array_to_storage_from_string_calldata_to_string(keccak256(pos_2, add(sub(_9, pos_2), 32)), value0, value1)
                        let pos_3 := mload(_2)
                        calldatacopy(pos_3, value0, value1)
                        let _10 := add(pos_3, value1)
                        mstore(_10, _3)
                        copy_byte_array_to_storage_from_string_calldata_to_string(keccak256(pos_3, add(sub(_10, pos_3), 32)), value4, value5)
                        let memPtr_1 := mload(_2)
                        finalize_allocation_6438(memPtr_1)
                        mstore(memPtr_1, /** @src 0:1248:1258  "msg.sender" */ caller())
                        /// @src 0:58:2853  "contract State {..."
                        mstore(/** @src 0:1243:1284  "User(msg.sender, identity, peer_id, addr)" */ add(memPtr_1, /** @src 0:58:2853  "contract State {..." */ 32), abi_decode_available_length_string(/** @src 0:1243:1284  "User(msg.sender, identity, peer_id, addr)" */ value0, value1, /** @src 0:58:2853  "contract State {..." */ calldatasize()))
                        mstore(/** @src 0:1243:1284  "User(msg.sender, identity, peer_id, addr)" */ add(memPtr_1, /** @src 0:58:2853  "contract State {..." */ _2), abi_decode_available_length_string(/** @src 0:1243:1284  "User(msg.sender, identity, peer_id, addr)" */ value2, value3, /** @src 0:58:2853  "contract State {..." */ calldatasize()))
                        mstore(/** @src 0:1243:1284  "User(msg.sender, identity, peer_id, addr)" */ add(memPtr_1, /** @src 0:58:2853  "contract State {..." */ 96), abi_decode_available_length_string(/** @src 0:1243:1284  "User(msg.sender, identity, peer_id, addr)" */ value4, value5, /** @src 0:58:2853  "contract State {..." */ calldatasize()))
                        mstore(_4, /** @src 0:1248:1258  "msg.sender" */ caller())
                        /// @src 0:58:2853  "contract State {..."
                        mstore(32, /** @src 0:1216:1228  "ownerMapping" */ 0x05)
                        /// @src 0:58:2853  "contract State {..."
                        copy_struct_to_storage_from_struct_User_to_struct_User(keccak256(_4, _2), memPtr_1)
                        /// @src 0:1306:1310  "addr"
                        fun_modDhtList(value4, value5)
                        /// @src 0:58:2853  "contract State {..."
                        return(_4, _4)
                    }
                    case 0x4ffe2a8b {
                        if callvalue() { revert(_4, _4) }
                        if slt(add(calldatasize(), not(3)), _4) { revert(_4, _4) }
                        let ret := fun_verification()
                        let memPos_3 := mload(_2)
                        mstore(memPos_3, iszero(iszero(ret)))
                        return(memPos_3, 32)
                    }
                    case 0x7228c6e0 {
                        if callvalue() { revert(_4, _4) }
                        let _11 := abi_decode_string(calldatasize())
                        let pos_4 := mload(_2)
                        let length_3 := mload(_11)
                        copy_memory_to_memory_with_cleanup(add(_11, 0x20), pos_4, length_3)
                        let end_1 := add(pos_4, length_3)
                        mstore(end_1, /** @src 0:2232:2247  "identityMapping" */ 0x03)
                        /// @src 0:58:2853  "contract State {..."
                        let converted_1 := copy_array_from_storage_to_memory_string(keccak256(pos_4, add(sub(end_1, pos_4), 0x20)))
                        let memPos_4 := mload(_2)
                        mstore(memPos_4, 0x20)
                        return(memPos_4, sub(abi_encode_string(converted_1, add(memPos_4, 0x20)), memPos_4))
                    }
                    case 0xcf1c23c1 {
                        if slt(add(calldatasize(), not(3)), 32) { revert(_4, _4) }
                        let offset_3 := calldataload(_3)
                        if gt(offset_3, 0xffffffffffffffff) { revert(_4, _4) }
                        let value0_1, value1_1 := abi_decode_string_calldata(add(_3, offset_3), calldatasize())
                        fun_modDhtList(value0_1, value1_1)
                        return(_4, _4)
                    }
                }
                revert(0, 0)
            }
            function copy_memory_to_memory_with_cleanup(src, dst, length)
            {
                let i := 0
                for { } lt(i, length) { i := add(i, 32) }
                {
                    mstore(add(dst, i), mload(add(src, i)))
                }
                mstore(add(dst, length), 0)
            }
            function abi_encode_string(value, pos) -> end
            {
                let length := mload(value)
                mstore(pos, length)
                copy_memory_to_memory_with_cleanup(add(value, 0x20), add(pos, 0x20), length)
                end := add(add(pos, and(add(length, 31), not(31))), 0x20)
            }
            function extract_byte_array_length(data) -> length
            {
                length := shr(1, data)
                let outOfPlaceEncoding := and(data, 1)
                if iszero(outOfPlaceEncoding) { length := and(length, 0x7f) }
                if eq(outOfPlaceEncoding, lt(length, 32))
                {
                    mstore(0, shl(224, 0x4e487b71))
                    mstore(4, 0x22)
                    revert(0, 0x24)
                }
            }
            function finalize_allocation_6438(memPtr)
            {
                let newFreePtr := add(memPtr, 128)
                if or(gt(newFreePtr, 0xffffffffffffffff), lt(newFreePtr, memPtr))
                {
                    mstore(0, shl(224, 0x4e487b71))
                    mstore(4, 0x41)
                    revert(0, 0x24)
                }
                mstore(64, newFreePtr)
            }
            function finalize_allocation_10236(memPtr)
            {
                let newFreePtr := add(memPtr, 64)
                if or(gt(newFreePtr, 0xffffffffffffffff), lt(newFreePtr, memPtr))
                {
                    mstore(0, shl(224, 0x4e487b71))
                    mstore(4, 0x41)
                    revert(0, 0x24)
                }
                mstore(64, newFreePtr)
            }
            function finalize_allocation(memPtr, size)
            {
                let newFreePtr := add(memPtr, and(add(size, 31), not(31)))
                if or(gt(newFreePtr, 0xffffffffffffffff), lt(newFreePtr, memPtr))
                {
                    mstore(0, shl(224, 0x4e487b71))
                    mstore(4, 0x41)
                    revert(0, 0x24)
                }
                mstore(64, newFreePtr)
            }
            function copy_array_from_storage_to_memory_string(slot) -> memPtr
            {
                memPtr := mload(64)
                let ret := /** @src -1:-1:-1 */ 0
                /// @src 0:58:2853  "contract State {..."
                let slotValue := sload(slot)
                let length := extract_byte_array_length(slotValue)
                mstore(memPtr, length)
                let _1 := 0x20
                let _2 := 1
                switch and(slotValue, _2)
                case 0 {
                    mstore(add(memPtr, _1), and(slotValue, not(255)))
                    ret := add(add(memPtr, shl(5, iszero(iszero(length)))), _1)
                }
                case 1 {
                    mstore(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ slot)
                    let dataPos := keccak256(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ _1)
                    let i := /** @src -1:-1:-1 */ 0
                    /// @src 0:58:2853  "contract State {..."
                    for { } lt(i, length) { i := add(i, _1) }
                    {
                        mstore(add(add(memPtr, i), _1), sload(dataPos))
                        dataPos := add(dataPos, _2)
                    }
                    ret := add(add(memPtr, i), _1)
                }
                finalize_allocation(memPtr, sub(ret, memPtr))
            }
            function abi_decode_available_length_string(src, length, end) -> array
            {
                if gt(length, 0xffffffffffffffff)
                {
                    mstore(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ shl(224, 0x4e487b71))
                    mstore(4, 0x41)
                    revert(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ 0x24)
                }
                let memPtr := mload(64)
                finalize_allocation(memPtr, add(and(add(length, 31), not(31)), 0x20))
                array := memPtr
                mstore(memPtr, length)
                if gt(add(src, length), end)
                {
                    revert(/** @src -1:-1:-1 */ 0, 0)
                }
                /// @src 0:58:2853  "contract State {..."
                calldatacopy(add(memPtr, 0x20), src, length)
                mstore(add(add(memPtr, length), 0x20), /** @src -1:-1:-1 */ 0)
            }
            /// @src 0:58:2853  "contract State {..."
            function abi_decode_string(dataEnd) -> value0
            {
                if slt(add(dataEnd, not(3)), 32) { revert(0, 0) }
                let offset := calldataload(4)
                if gt(offset, 0xffffffffffffffff) { revert(0, 0) }
                if iszero(slt(add(offset, 35), dataEnd))
                {
                    revert(/** @src -1:-1:-1 */ 0, 0)
                }
                /// @src 0:58:2853  "contract State {..."
                value0 := abi_decode_available_length_string(add(offset, 36), calldataload(add(4, offset)), dataEnd)
            }
            function abi_decode_string_calldata(offset, end) -> arrayPos, length
            {
                if iszero(slt(add(offset, 0x1f), end)) { revert(0, 0) }
                length := calldataload(offset)
                if gt(length, 0xffffffffffffffff) { revert(0, 0) }
                arrayPos := add(offset, 0x20)
                if gt(add(add(offset, length), 0x20), end) { revert(0, 0) }
            }
            function clear_storage_range_bytes1(start, end)
            {
                for { } lt(start, end) { start := add(start, 1) }
                { sstore(start, 0) }
            }
            function clean_up_bytearray_end_slots_string_storage(array, len, startIndex)
            {
                if gt(len, 31)
                {
                    mstore(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ array)
                    let data := keccak256(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ 0x20)
                    let deleteStart := add(data, shr(5, add(startIndex, 31)))
                    if lt(startIndex, 0x20) { deleteStart := data }
                    clear_storage_range_bytes1(deleteStart, add(data, shr(5, add(len, 31))))
                }
            }
            function extract_used_part_and_set_length_of_short_byte_array(data, len) -> used
            {
                used := or(and(data, not(shr(shl(3, len), not(0)))), shl(1, len))
            }
            function copy_struct_to_storage_from_struct_User_to_struct_User(slot, value)
            {
                sstore(slot, or(and(sload(slot), shl(160, 0xffffffffffffffffffffffff)), and(mload(value), sub(shl(160, 1), 1))))
                let _1 := 1
                let memberSlot := add(slot, _1)
                let _2 := 32
                let _3 := mload(add(value, _2))
                let newLen := mload(_3)
                let _4 := 0xffffffffffffffff
                if gt(newLen, _4)
                {
                    mstore(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ shl(224, 0x4e487b71))
                    mstore(4, 0x41)
                    revert(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ 0x24)
                }
                clean_up_bytearray_end_slots_string_storage(memberSlot, extract_byte_array_length(sload(memberSlot)), newLen)
                let srcOffset := /** @src -1:-1:-1 */ 0
                /// @src 0:58:2853  "contract State {..."
                srcOffset := _2
                switch gt(newLen, 31)
                case 1 {
                    let loopEnd := and(newLen, not(31))
                    mstore(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ memberSlot)
                    let dstPtr := keccak256(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ _2)
                    let i := /** @src -1:-1:-1 */ 0
                    /// @src 0:58:2853  "contract State {..."
                    for { } lt(i, loopEnd) { i := add(i, _2) }
                    {
                        sstore(dstPtr, mload(add(_3, srcOffset)))
                        dstPtr := add(dstPtr, _1)
                        srcOffset := add(srcOffset, _2)
                    }
                    if lt(loopEnd, newLen)
                    {
                        let lastValue := mload(add(_3, srcOffset))
                        sstore(dstPtr, and(lastValue, not(shr(and(shl(3, newLen), 248), not(0)))))
                    }
                    sstore(memberSlot, add(shl(_1, newLen), _1))
                }
                default {
                    let value_1 := /** @src -1:-1:-1 */ 0
                    /// @src 0:58:2853  "contract State {..."
                    if newLen
                    {
                        value_1 := mload(add(_3, srcOffset))
                    }
                    sstore(memberSlot, extract_used_part_and_set_length_of_short_byte_array(value_1, newLen))
                }
                let memberSlot_1 := add(slot, 2)
                let _5 := mload(add(value, 64))
                let newLen_1 := mload(_5)
                if gt(newLen_1, _4)
                {
                    mstore(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ shl(224, 0x4e487b71))
                    mstore(4, 0x41)
                    revert(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ 0x24)
                }
                clean_up_bytearray_end_slots_string_storage(memberSlot_1, extract_byte_array_length(sload(memberSlot_1)), newLen_1)
                let srcOffset_1 := /** @src -1:-1:-1 */ 0
                /// @src 0:58:2853  "contract State {..."
                srcOffset_1 := _2
                switch gt(newLen_1, 31)
                case 1 {
                    let loopEnd_1 := and(newLen_1, not(31))
                    mstore(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ memberSlot_1)
                    let dstPtr_1 := keccak256(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ _2)
                    let i_1 := /** @src -1:-1:-1 */ 0
                    /// @src 0:58:2853  "contract State {..."
                    for { } lt(i_1, loopEnd_1) { i_1 := add(i_1, _2) }
                    {
                        sstore(dstPtr_1, mload(add(_5, srcOffset_1)))
                        dstPtr_1 := add(dstPtr_1, _1)
                        srcOffset_1 := add(srcOffset_1, _2)
                    }
                    if lt(loopEnd_1, newLen_1)
                    {
                        let lastValue_1 := mload(add(_5, srcOffset_1))
                        sstore(dstPtr_1, and(lastValue_1, not(shr(and(shl(3, newLen_1), 248), not(0)))))
                    }
                    sstore(memberSlot_1, add(shl(_1, newLen_1), _1))
                }
                default {
                    let value_2 := /** @src -1:-1:-1 */ 0
                    /// @src 0:58:2853  "contract State {..."
                    if newLen_1
                    {
                        value_2 := mload(add(_5, srcOffset_1))
                    }
                    sstore(memberSlot_1, extract_used_part_and_set_length_of_short_byte_array(value_2, newLen_1))
                }
                let memberSlot_2 := add(slot, 3)
                let _6 := mload(add(value, 96))
                let newLen_2 := mload(_6)
                if gt(newLen_2, _4)
                {
                    mstore(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ shl(224, 0x4e487b71))
                    mstore(4, 0x41)
                    revert(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ 0x24)
                }
                clean_up_bytearray_end_slots_string_storage(memberSlot_2, extract_byte_array_length(sload(memberSlot_2)), newLen_2)
                let srcOffset_2 := /** @src -1:-1:-1 */ 0
                /// @src 0:58:2853  "contract State {..."
                srcOffset_2 := _2
                switch gt(newLen_2, 31)
                case 1 {
                    let loopEnd_2 := and(newLen_2, not(31))
                    mstore(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ memberSlot_2)
                    let dstPtr_2 := keccak256(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ _2)
                    let i_2 := /** @src -1:-1:-1 */ 0
                    /// @src 0:58:2853  "contract State {..."
                    for { } lt(i_2, loopEnd_2) { i_2 := add(i_2, _2) }
                    {
                        sstore(dstPtr_2, mload(add(_6, srcOffset_2)))
                        dstPtr_2 := add(dstPtr_2, _1)
                        srcOffset_2 := add(srcOffset_2, _2)
                    }
                    if lt(loopEnd_2, newLen_2)
                    {
                        let lastValue_2 := mload(add(_6, srcOffset_2))
                        sstore(dstPtr_2, and(lastValue_2, not(shr(and(shl(3, newLen_2), 248), not(0)))))
                    }
                    sstore(memberSlot_2, add(shl(_1, newLen_2), _1))
                }
                default {
                    let value_3 := /** @src -1:-1:-1 */ 0
                    /// @src 0:58:2853  "contract State {..."
                    if newLen_2
                    {
                        value_3 := mload(add(_6, srcOffset_2))
                    }
                    sstore(memberSlot_2, extract_used_part_and_set_length_of_short_byte_array(value_3, newLen_2))
                }
            }
            function copy_byte_array_to_storage_from_string_calldata_to_string(slot, src, len)
            {
                if gt(len, 0xffffffffffffffff)
                {
                    mstore(0, shl(224, 0x4e487b71))
                    mstore(4, 0x41)
                    revert(0, 0x24)
                }
                clean_up_bytearray_end_slots_string_storage(slot, extract_byte_array_length(sload(slot)), len)
                let srcOffset := 0
                switch gt(len, 31)
                case 1 {
                    let loopEnd := and(len, not(31))
                    mstore(srcOffset, slot)
                    let _1 := 0x20
                    let dstPtr := keccak256(srcOffset, _1)
                    let i := srcOffset
                    for { } lt(i, loopEnd) { i := add(i, _1) }
                    {
                        sstore(dstPtr, calldataload(add(src, srcOffset)))
                        dstPtr := add(dstPtr, 1)
                        srcOffset := add(srcOffset, _1)
                    }
                    if lt(loopEnd, len)
                    {
                        sstore(dstPtr, and(calldataload(add(src, srcOffset)), not(shr(and(shl(3, len), 248), not(0)))))
                    }
                    sstore(slot, add(shl(1, len), 1))
                }
                default {
                    let value := 0
                    if len
                    {
                        value := calldataload(add(src, srcOffset))
                    }
                    sstore(slot, extract_used_part_and_set_length_of_short_byte_array(value, len))
                }
            }
            /// @ast-id 120 @src 0:1492:1724  "function verification() public view returns(bool) {..."
            function fun_verification() -> var
            {
                /// @src 0:58:2853  "contract State {..."
                mstore(/** @src -1:-1:-1 */ 0, /** @src 0:1596:1606  "msg.sender" */ caller())
                /// @src 0:58:2853  "contract State {..."
                mstore(0x20, /** @src 0:1583:1595  "ownerMapping" */ 0x05)
                /// @src 0:58:2853  "contract State {..."
                let _1 := sload(keccak256(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ 0x40))
                /// @src 0:1566:1614  "abi.encodePacked(ownerMapping[msg.sender].owner)"
                let expr_103_mpos := /** @src 0:58:2853  "contract State {..." */ mload(0x40)
                /// @src 0:1566:1614  "abi.encodePacked(ownerMapping[msg.sender].owner)"
                let _2 := add(expr_103_mpos, /** @src 0:58:2853  "contract State {..." */ 0x20)
                mstore(_2, and(shl(96, _1), not(0xffffffffffffffffffffffff)))
                /// @src 0:1566:1614  "abi.encodePacked(ownerMapping[msg.sender].owner)"
                mstore(expr_103_mpos, 20)
                finalize_allocation_10236(expr_103_mpos)
                /// @src 0:1556:1615  "keccak256(abi.encodePacked(ownerMapping[msg.sender].owner))"
                let expr := keccak256(/** @src 0:58:2853  "contract State {..." */ _2, mload(/** @src 0:1556:1615  "keccak256(abi.encodePacked(ownerMapping[msg.sender].owner))" */ expr_103_mpos))
                /// @src 0:1629:1657  "abi.encodePacked(msg.sender)"
                let expr_mpos := /** @src 0:58:2853  "contract State {..." */ mload(0x40)
                /// @src 0:1629:1657  "abi.encodePacked(msg.sender)"
                let _3 := add(expr_mpos, /** @src 0:58:2853  "contract State {..." */ 0x20)
                mstore(_3, shl(96, /** @src 0:1596:1606  "msg.sender" */ caller()))
                /// @src 0:1629:1657  "abi.encodePacked(msg.sender)"
                mstore(expr_mpos, /** @src 0:1566:1614  "abi.encodePacked(ownerMapping[msg.sender].owner)" */ 20)
                /// @src 0:1629:1657  "abi.encodePacked(msg.sender)"
                finalize_allocation_10236(expr_mpos)
                /// @src 0:1552:1696  "if (keccak256(abi.encodePacked(ownerMapping[msg.sender].owner)) == keccak256(abi.encodePacked(msg.sender))) {..."
                if /** @src 0:1556:1658  "keccak256(abi.encodePacked(ownerMapping[msg.sender].owner)) == keccak256(abi.encodePacked(msg.sender))" */ eq(expr, /** @src 0:1619:1658  "keccak256(abi.encodePacked(msg.sender))" */ keccak256(/** @src 0:58:2853  "contract State {..." */ _3, mload(/** @src 0:1619:1658  "keccak256(abi.encodePacked(msg.sender))" */ expr_mpos)))
                /// @src 0:1552:1696  "if (keccak256(abi.encodePacked(ownerMapping[msg.sender].owner)) == keccak256(abi.encodePacked(msg.sender))) {..."
                {
                    /// @src 0:1674:1685  "return true"
                    var := /** @src 0:1681:1685  "true" */ 0x01
                    /// @src 0:1674:1685  "return true"
                    leave
                }
                /// @src 0:1705:1717  "return false"
                var := /** @src -1:-1:-1 */ 0
            }
            /// @src 0:58:2853  "contract State {..."
            function array_push_from_string_calldata_to_array_string_storage_dyn_ptr(value0, value1)
            {
                /// @src 0:2737:2744  "dhtList"
                let _1 := 0x00
                /// @src 0:58:2853  "contract State {..."
                let oldLen := sload(/** @src 0:2737:2744  "dhtList" */ _1)
                /// @src 0:58:2853  "contract State {..."
                if iszero(lt(oldLen, 18446744073709551616))
                {
                    mstore(/** @src 0:2737:2744  "dhtList" */ _1, /** @src 0:58:2853  "contract State {..." */ shl(224, 0x4e487b71))
                    mstore(4, 0x41)
                    revert(/** @src 0:2737:2744  "dhtList" */ _1, /** @src 0:58:2853  "contract State {..." */ 0x24)
                }
                let _2 := add(oldLen, 1)
                sstore(/** @src 0:2737:2744  "dhtList" */ _1, /** @src 0:58:2853  "contract State {..." */ _2)
                if iszero(lt(oldLen, _2))
                {
                    mstore(/** @src 0:2737:2744  "dhtList" */ _1, /** @src 0:58:2853  "contract State {..." */ shl(224, 0x4e487b71))
                    mstore(4, 0x32)
                    revert(/** @src 0:2737:2744  "dhtList" */ _1, /** @src 0:58:2853  "contract State {..." */ 0x24)
                }
                mstore(/** @src 0:2737:2744  "dhtList" */ _1, _1)
                /// @src 0:58:2853  "contract State {..."
                copy_byte_array_to_storage_from_string_calldata_to_string(add(18569430475105882587588266137607568536673111973893317399460219858819262702947, oldLen), value0, value1)
            }
            /// @ast-id 208 @src 0:2440:2851  "function modDhtList(string calldata addr) public payable {..."
            function fun_modDhtList(var_addr_offset, var_addr_length)
            {
                /// @src 0:2541:2551  "dhtListTtl"
                let _1 := 0x01
                /// @src 0:58:2853  "contract State {..."
                let diff := sub(/** @src 0:2523:2538  "block.timestamp" */ timestamp(), /** @src 0:58:2853  "contract State {..." */ sload(/** @src 0:2541:2551  "dhtListTtl" */ _1))
                /// @src 0:58:2853  "contract State {..."
                if gt(diff, /** @src 0:2523:2538  "block.timestamp" */ timestamp())
                /// @src 0:58:2853  "contract State {..."
                {
                    mstore(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ shl(224, 0x4e487b71))
                    mstore(4, 0x11)
                    revert(/** @src -1:-1:-1 */ 0, /** @src 0:58:2853  "contract State {..." */ 0x24)
                }
                /// @src 0:2562:2845  "if (checkTtl > 60) {..."
                switch /** @src 0:2566:2579  "checkTtl > 60" */ gt(diff, /** @src 0:2577:2579  "60" */ 0x3c)
                case /** @src 0:2562:2845  "if (checkTtl > 60) {..." */ 0 {
                    /// @src -1:-1:-1
                    let _2 := 0
                    /// @src 0:2737:2751  "dhtList.length"
                    let expr := /** @src 0:58:2853  "contract State {..." */ sload(/** @src -1:-1:-1 */ _2)
                    /// @src 0:2733:2803  "if (dhtList.length > 2) {..."
                    if /** @src 0:2737:2755  "dhtList.length > 2" */ gt(expr, /** @src 0:2754:2755  "2" */ 0x02)
                    /// @src 0:2733:2803  "if (dhtList.length > 2) {..."
                    {
                        /// @src 0:58:2853  "contract State {..."
                        if iszero(expr)
                        {
                            mstore(/** @src -1:-1:-1 */ _2, /** @src 0:58:2853  "contract State {..." */ shl(224, 0x4e487b71))
                            mstore(4, 0x31)
                            revert(/** @src -1:-1:-1 */ _2, /** @src 0:58:2853  "contract State {..." */ 0x24)
                        }
                        let newLen := add(expr, not(0))
                        if iszero(lt(newLen, expr))
                        {
                            mstore(/** @src -1:-1:-1 */ _2, /** @src 0:58:2853  "contract State {..." */ shl(224, 0x4e487b71))
                            mstore(4, 0x32)
                            revert(/** @src -1:-1:-1 */ _2, /** @src 0:58:2853  "contract State {..." */ 0x24)
                        }
                        mstore(/** @src -1:-1:-1 */ _2, _2)
                        /// @src 0:58:2853  "contract State {..."
                        let slot := add(expr, 0x290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e562)
                        let oldLen := extract_byte_array_length(sload(slot))
                        if iszero(iszero(oldLen))
                        {
                            switch gt(oldLen, 31)
                            case 1 {
                                mstore(/** @src -1:-1:-1 */ _2, /** @src 0:58:2853  "contract State {..." */ slot)
                                let data := keccak256(/** @src -1:-1:-1 */ _2, /** @src 0:58:2853  "contract State {..." */ 32)
                                clear_storage_range_bytes1(add(data, /** @src 0:2541:2551  "dhtListTtl" */ _1), /** @src 0:58:2853  "contract State {..." */ add(data, shr(5, add(oldLen, 31))))
                                let data_1 := keccak256(/** @src -1:-1:-1 */ _2, /** @src 0:58:2853  "contract State {..." */ 32)
                                sstore(slot, /** @src -1:-1:-1 */ _2)
                                /// @src 0:58:2853  "contract State {..."
                                sstore(data_1, /** @src -1:-1:-1 */ _2)
                            }
                            default /// @src 0:58:2853  "contract State {..."
                            {
                                sstore(slot, /** @src -1:-1:-1 */ _2)
                            }
                        }
                        /// @src 0:58:2853  "contract State {..."
                        sstore(/** @src -1:-1:-1 */ _2, /** @src 0:58:2853  "contract State {..." */ newLen)
                    }
                    /// @src 0:2816:2834  "dhtList.push(addr)"
                    array_push_from_string_calldata_to_array_string_storage_dyn_ptr(var_addr_offset, var_addr_length)
                }
                default /// @src 0:2562:2845  "if (checkTtl > 60) {..."
                {
                    /// @src 0:58:2853  "contract State {..."
                    let memPtr := mload(64)
                    let _3 := 0x20
                    let newFreePtr := add(memPtr, _3)
                    if or(gt(newFreePtr, 0xffffffffffffffff), lt(newFreePtr, memPtr))
                    {
                        mstore(0, shl(224, 0x4e487b71))
                        mstore(4, 0x41)
                        revert(0, 0x24)
                    }
                    mstore(64, newFreePtr)
                    /// @src -1:-1:-1
                    let _4 := 0
                    /// @src 0:58:2853  "contract State {..."
                    mstore(memPtr, /** @src -1:-1:-1 */ _4)
                    /// @src 0:58:2853  "contract State {..."
                    let oldLen_1 := sload(/** @src -1:-1:-1 */ _4)
                    /// @src 0:58:2853  "contract State {..."
                    sstore(/** @src -1:-1:-1 */ _4, _4)
                    /// @src 0:58:2853  "contract State {..."
                    if iszero(iszero(oldLen_1))
                    {
                        mstore(/** @src -1:-1:-1 */ _4, _4)
                        /// @src 0:58:2853  "contract State {..."
                        let _5 := 18569430475105882587588266137607568536673111973893317399460219858819262702947
                        let _6 := add(_5, oldLen_1)
                        let start := _5
                        for { }
                        lt(start, _6)
                        {
                            start := add(start, /** @src 0:2541:2551  "dhtListTtl" */ _1)
                        }
                        /// @src 0:58:2853  "contract State {..."
                        {
                            let oldLen_2 := extract_byte_array_length(sload(start))
                            if iszero(iszero(oldLen_2))
                            {
                                let _7 := 31
                                switch gt(oldLen_2, _7)
                                case 1 {
                                    mstore(/** @src -1:-1:-1 */ _4, /** @src 0:58:2853  "contract State {..." */ start)
                                    let data_2 := keccak256(/** @src -1:-1:-1 */ _4, /** @src 0:58:2853  "contract State {..." */ _3)
                                    clear_storage_range_bytes1(add(data_2, /** @src 0:2541:2551  "dhtListTtl" */ _1), /** @src 0:58:2853  "contract State {..." */ add(data_2, shr(5, add(oldLen_2, _7))))
                                    let data_3 := keccak256(/** @src -1:-1:-1 */ _4, /** @src 0:58:2853  "contract State {..." */ _3)
                                    sstore(start, /** @src -1:-1:-1 */ _4)
                                    /// @src 0:58:2853  "contract State {..."
                                    sstore(data_3, /** @src -1:-1:-1 */ _4)
                                }
                                default /// @src 0:58:2853  "contract State {..."
                                {
                                    sstore(start, /** @src -1:-1:-1 */ _4)
                                }
                            }
                        }
                    }
                    /// @src 0:58:2853  "contract State {..."
                    mstore(/** @src -1:-1:-1 */ _4, _4)
                    /// @src 0:58:2853  "contract State {..."
                    sstore(/** @src 0:2541:2551  "dhtListTtl" */ _1, /** @src 0:2523:2538  "block.timestamp" */ timestamp())
                    /// @src 0:2676:2694  "dhtList.push(addr)"
                    array_push_from_string_calldata_to_array_string_storage_dyn_ptr(var_addr_offset, var_addr_length)
                }
            }
        }
        data ".metadata" hex"a2646970667358221220fd98970d04d660754a68ccedf59801d6a72fcd37f550f92c5c8ec225887934fb64736f6c63430008110033"
    }
}
