package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

const input = `4601
1583
2995
5319
3352
1722
4331
5840
3339
5341
3415
1297
1610
2703

5928
1266
6405
4354
2815
1621
3545
1733
2162
1155
3674
4055
4262
2115

25873
16103
17042

6479
1521
6902
6331
6146
1591
2063
2149
1463
1865
2598
6893
3617

1431
8935
8134
8885
8698
8884
7982
3617
7079

34031
8383

3697
1199
3703
1322
5679
1627
5910
1092
6445
2371
3672
2281
2710
5111

17480

58355

3711
2406
2101
1406
5376
2149
4608
6768
6322
3792
4249
2342
2014

6197
6908
8484
3901
3001
7880
5165
2536

3964
4229
4898
4544
3535
6329
1573
4988
1424

57606

16325

6038
11538
13026
12734
6077

4696
3914
2375
1188
1247
1594
1843
5297
6053
1114
3179
5968
1430
4268
2514

6684
2321
3946
1746
7583
10379
3696
10110

5834
5944
9336
9638
3373
4279
9942

7208
7801
3771
8183
1824
8217
1545
3041
7369
5980

8368
6180
3100
4801
4430
6758
3023
3394

2545
4992
3347
5830
3764
3237
2431
3406
1408
5601
6548
4625

32334

1112
5471
10983
8298
4342
4221
7807

5733
2012
5928
3503
2785
4867
1515
4819
3569
3353
5997
4921
5498
3234
5588

11040
5547
12070
2168
7603
7245
4966

13027
12518
15381
3251
14935

15858
29295

2476
4289
10350
3953
10089
6953
5555
3007

61732

5567
4895
6095
4880
2945
6022
2129
5010
4100
3554
3410
1820

9902
33605

15331
15615
5062
15073

5940
2252
1233
2582
3624
2754
2304
3189
4060
3862
5164
1487
2776
1807
3385

6798
2385
3755
10299
11500
6478
1911

1469
9309
1785
17145

3363
3210
3285
4491
5992
2927
4098
2133
3191
2810
4694

31006
1582

4356
6146
2051
6503
9487
9244
4808
7761
7108

1771
6062
7135
5575
3278
1221
2547
7266
5444
7395

10634
1196

1361

2810
3106
8654
9455
3778
4219
8882
9387
5128

7114
1441
6146
4752
2388
6420
7350
1624
6050
5469
6751

5570
4674

2074
3476
4064
9446
2889
4871
6298
2970
8588

3482
8492
2226
4951
6733
5463
3183
4807
4806

6974
10251
4637
9558
5153
4665

6860
6949
1524
4201
4397
2278
4909
1353
2074
1920
7439
1040

4135
1411
5536
8004
3090
1036
6613
6353
6810

3932
4743
1615
3201
3791
6357
2511
3440
2290
4741
2844
3121
3641
3497

4549
5142
4438
2945
4576
4871
4795
6007
4884
1031
2111
4388
4705
5725
3425

45444

6688
9460
2727
6989
6789
12460

7310
5910
7822
7627
7600
4793
3152
5747
5197
8006

23983
33879

1020
4783
5117
3628
3065
6131
5543
1576
1159
3721
3182
1409
2738
5084

42750

8509
4370
1042
5664

59445

1589
5133
6413
6578
3301
8256
7749
4321
8190
5025

27739
18609

5447
4022
4069
1437
2811
6929
2754
1108
7390
3995
3480
3225

4240
6959
3289
8063
4751
4896
4611
6256
1105
1334
7598

27831

1821
5333
7853
6574
7985
3972
2306
1875
1857
6386

12780
2513
22939

9625
3785
8050
7432
11772
10914

1387
3439
4644
3880
1878
3911
5329
1480
5779
2083
3886
5525
3930
4813
2570

15763
2839
14258
1460

6193
3397
3222
3262
3735
2406
5588
2068
2862
3271
1401
3957
5690
4632

15070
5164
8902
9580
4271

4494
5665
7242
7687
7031
6135
2670
3315
6146
8085

2625
24285

4414
5985
2681
7036
6486
2501
4408
8582
5150
4295

1082
1497
5638
4643
1413
5969
1515
4016
4856
3287
3122
2948
3711
5031
4692

6445
2261
2132
3517
2461
2767
2071
2176
6357
1304
6034
1346

1318
4719
2054
1433
6514
3729
4329
1026
1502
2611
1363
6379
3617

1885
7519
2849
5167
3136
4819
8018
4110
2666
6098
4943

6575
1624
6004
4178
7453
5979
6736
3883
7109
1271
1175

9494
12633
11576
3547

4653
6777
4497
2246
4666
4675
4440
7227
3808
2425
6014
7305

2538
6155
4321
6239
6279
1445
5592

4920
5435
1392
4122
10513
6359
1574
7222

9013
2150
8377
8884
4108
5427
1605
6907

7861
10554
1210
1346
5296
4419
9147

1858
2956
2070
10271
5810
2632
9429
3941

4306
4972
6216
6345
3971
3920
6295
3552
1399
5425
5218
4512
3376
4460

4107
4484
5148
5456
7902
2185
3439
7199
9514

11030
12096

2886
1561
14057
7480
12186

8600
4085
11755
4564
1697
6345
10176

2538
10138
8328
10621
13322
13907

5450
3403
5776
5967
3873
2761
5992
2251
2278
4673
3442
2720
2556
4848
5967

2806
8228
9047
10630
4367
4917
4084
7893

2842
18343
1497
19962

3985
10986
1431
1569
8053
8987
6474

5733
2237
6183
1642
1005
4389
1889
2385
2470
6338
2200
3399
3659
4092

6579
5872
4229
1081
2254
1229
5218
4267
6906
2818
1084
3655
2294

1882
5005
3805
2088
6809
4982
1935
6789
1961
1842
5988
2859
3965

2313
1021
5090
1309
5110
3142
4102
2919
6010
3079
5786
3701
2659
5728

4251
1899
4119
8941
6677
2161
2119
2772
3122

7821
6113
5150
3486
2072
4089
2820
3161
7589
1787
5988

10582
8240
1934
1504
9070
1186
7766
1066

6700
6181
3238
1940
1229
3688
1375
5873
2328
6384
7446
7246

5542
6771
1964
4532
3011
2300
1451
6726
4595
2926
5237
5020
6129

15008
4409
6769
2211

11825
8947
17411
2893

8460
11664
2304
6865
2132
9174

15540
16925
12015
7621

1439
11886
3028

4604
1057
5342
4723
7773
3373
7488
8028
3246
6355
1939

2921
4447
3472
5175
5909
5936
5263
3993
1589
3906
4865
3156
2010

10089

5964
7650
4692
1537
3281
3162
1064
4567
3910
2229
3566

15533
17311
11818
6179

4583
4577
3573
1379
3761
1048
4799
4275
3515
5637
5435
3801
5933
1926

11841
16030
3213

4307
7300
1237
5513
6799
2402
2554

67023

1573
6798
3002
7536
2373
5994
10122

1149
1836
7129
6621
7432
2468
4028
6560
5127
3781
7035
5845

13729
10212
10942
10087
10080
6493

6850
5969
3370
5400
6433
4762
5702
6193
1836
1507
2637
3051
1394

5465
10565
4923
1671
9789
9150
6466
8013

31283

2001
6032
6765
3896
7593
3175
3447
1718

6062
7630
6302
4448
1014
4756
4696
4069
5663
2181
6153

9228
1871
11739
10894

7435
6290
3249
9358
9196
8951
5023
4910
9686

3823
4247
7159

4630
3352
7700
7752
3343
4577
2448
10326

10002
28943

2592
4430
4077
2126
4628
6466
5919
7019
2667
2573
7038
3153

2837
12151
9774
1574
3046
13079

6949

5306
7432
4274
6835
7057
5695
4727
5581
6516
3903
3138

4083
4247
1806
2753
3612
3509
5498
2160
5814
1502
5760
4663
1354
1592

59040

3738
7624
4539
8649
9418
5495
1178
8389

4458
5283
2572
3530
6228
4711
2621
2131
6363
5611
3236
2314
2547
1905

3744
1413
2997

2214
4673
4481
5827
3752
5489
5301
7568
6869
2928
2900

11668
7579
2324
7245
14362

1191
5622
2764
3241
3690
3627
2637
4043
3012
2773
4276
1127
3649

5327
6535
4821
5531
3037
4700
6287
6476
2266
7550

3054
5023
1296
4394
5649
1822
6609
7038
5836
6571
6918

9473
5377
8449
4839
3118
7861

8784
3880
5133
6131
2114
5196
10405
1040

5003
5888
6058
6910
3155
3278
1159
3863
4565
1181
2716
6431
3014

7633
4984
1297
4524
5183
3717
12038

5908
7784
11211
7635
2060
5502
9179

6537
6895
3455
4799
5557
3497
4467
1694
5227
1663
2377
3538
4529

9371
2421
9789
7373
6459
6405
7775
8218

2077
8589
3417
7811
10255
4325
7168

5963
2806
3512
2369
5721
3852
3022
3764
6152
1158
6460
4694
1909
2655

7789
3030
3621
6269
3149
6405
2288
3512
5674
1301

6120
2629
6348
7105
3459
6678
5903
4519
3074
2086
3927

2924
12290
1873
5613
3098
8641

24464
37206

7180
13271
4832
10691
9084

4915
1440
1841
1476
7124
1692
7480
4583
1896
3475
5424

6864
3187
4680
5159
4314
3148
1060
7290
2080
8308

2412
16405
15203
5606
9411

13201
6242
11906
15782
12558

5319
7968
6732
9702

7203
3571
7241
8764
3466
6032
6608
6898
7189
2749

6184
7590
4976
4289
8553
8238
5844
8314

3264
2924
2586
1234
3126
5269
6147
6029
2384
6440
6161
1708
2928
2336

24565
31964

11118
10621
1430
3312
1135
6740
9456

16874
4385
8675
4237

1537
4313
5552
2195
3530
2433
3056
3693
1151
4297
2017
4917
3823
3388
1580

5795
2642
3822
1824
2366
3800
2020
4065
2758
2677
3951
5175
4433
1199
3212

4183
2822
6355
4427
6384
2931
4892
3148
5802
3885
4892
5233
4269
4286

3711
5430
6115
5561
5066
2796
3886
1960
6083
4657

11662
12901
6891
3897

1086
4333
1477
1046
5670
3996
1909
3544
5143
1387
6319
3084
3136
6084

36111
14790

6552
8491

6636
9718
15736
1619

9531
7507
7311
5991
6901
12139
10580

10956
7878
9100
11884
2906
10388

7767
4040
5872
1060
2299
1008
3597
2749
5025

1550

5364
7108
8661
8894
9240
10008
4243

46695

4195
2850
2978
4995
2944
5773
3669
3662
4471
6343
4258
1252
6408
5995

9482
11898
9022
8167
9232

2647
1990
2948
5314
2341
3828
2513
2091
4173

9640
2690
8774
11619
11201
2429
7216

3433
5914
5223
4879
1349
4692
2906
3761
4416
4813
2751
4474
2057
1675

9803
8986
3617
2175
1412
8370

7671
6448
6759
1857

25790
14495

1872
6441
4585
4053
4142
5844
2089
3387
2660
5890
6499
3769
2710

23432
16915

1351
2656
6557
1149
5054
1681
6871
2643
3158
1617
4087
6498
4993

5494
10055
2730
5800
2923
8037
3448
1527

5803
18165
21951

4766
3809
4678
3184
9687
7825
7222
6638
2016

1801
1717
4955
6088
1848
4033
7023
3551
5970
1415
7027
7411

25450
11523
22548

1566
4742
3947
2708
2019
3479
4735
6002
7042
1828
7071
1535

14218
5232
10228
3340
7975

8837
3510
10363
11854
2821
7627
10186

1102
3749
1436
1561
3098
3849
3726
2041
1100
1532
5473
4643
5922
4739
1216

4408
7369
1369
4742
7109
4368
3468
2777
3296
6999
5777
5019

3750
12190

15694
12690
14834
15532
1955

8702
1209
3012
3074
8390
6779
6117
2544
3900
6374

4126
5772
7149
3949
2371
2374
2469
7839
7898
2871
5808

1108

9610
4374
10275
2595
9413
8906
6402
1551

6560
8632
5064
2152
3443
4229

13104
24831
17968

4648
5332
4866
1048
4089
4710
1406
2544
5319
4685
4622
6080
5899
3266

5694
12879
4015
16066
1600

11186
31873

3871
6382
5076
3606
4283
7913
11124

9959
14807
17825
5880

3914
14183

8104
8322
10254
3957
1410
9402
2012

1130
4640
5738
3860
1510
3067
5021
5321
2820
1269
3433
5216
5585
1344

5975
9879
4969
9796
1404
6772
3240
1954

12512
13741
6425
13234
10597
12414

2031
2255
3306
1831
1601
1959
4452
5054
4967
5392
4040
4989
4385
5838
5043

3165
10185
10580
1626
5436
9644
10842

5139
8090
1045
5867
5987
8488
2409

1045
2843
3543
3224
2855
3318
4192
4600
5773
4074
4705
4716
1798
5274
2443

6325
7878
4932
7543
8073
2113
5481
8138
2151
1296

6007
2506
2346
3631
1526
4875
3099
5253
6281
2057
3209
2622
2930
4926

2901
4561
2361
4474
8508
6515
5024
2201
4402
6922

4219
24714
17196

38164

1900
5727
3935
5498
6165
4621
3347
4387
5982
4213
1742
1302
2452
3923

9558
7131
15662
9101
13730

6474
2222
6539
3747
5995
5573
6288
4684
5121
2226
1722
1117
1020

3565
3489

18395
20407

57158

19992

3636
4392
4584
5421
3159
5817
2303
1013
1758
1802
4630
4261
1524
1639

4652
1967
4765
4946
3023
2732
2749
5258
5935
6428
3053
6328
2389
5681

5450
1993
2824
3596
2658
8545
6577
3024
4001
1780

7902
6187

1349
6028
3307
1964
5116
2308
6108
4286
1689
3324
3728
4946
4267
3801
3806

14183
29320

32460
25012

3317
5545
7302
4936
7934
6094
7971
2497
8130
3752

4195
2147
7770
4214
4737
7962
5942
4478
3816
7508
6693

26013
8614
19828

1144
2361
4371
4246
3442
1693
5137
5806
2434
1413
3135
4028
4757
6005

2351
4927
3004
3384
1545
2172
3407
4120
3954
1806
2342
5847
5228
1005

11490
7855
9775
10577
3232
5211

8316
2984
13490
15716
7433

11288
14551
10056
14839
7241

4573
2731
3358
4850
6352
5419
1021
6611
6489
3794
6089`

// heap is a generic implementation of a min heap.
type heap[T constraints.Ordered] struct {
	values []T
}

// push inserts an element into the heap.
func (h *heap[T]) push(x T) {
	h.values = append(h.values, x)
	h.siftDown(0, len(h.values)-1)
}

// pop removes the min element from the heap.
func (h *heap[T]) pop() (*T, bool) {
	if len(h.values) == 0 {
		return nil, false
	}
	last := h.values[len(h.values)-1]
	h.values = h.values[:len(h.values)-1]
	if len(h.values) == 0 {
		return &last, true
	}
	first := h.values[0]
	h.values[0] = last
	h.siftUp(0)
	return &first, true
}

// siftDown moves an element inserted at index j (a leaf) up in the tree
// (i.e. down in the ascending order of the heap) to maintain the heap invariant.
func (h *heap[T]) siftDown(i, j int) {
	item := h.values[j]
	for i < j {
		parentIndex := (j - 1) >> 1
		parent := h.values[parentIndex]
		if item < parent {
			h.values[j] = parent
			j = parentIndex
			continue
		}
		break
	}
	h.values[j] = item
}

// siftUp moves an element inserted at "index" down in the tree
// (i.e. up in the ascending order of the heap) to maintain the heap invariant.
func (h *heap[T]) siftUp(index int) {
	// Move item to a leaf node, swapping with the smaller of the two
	// kids nodes on the way there.
	originalIndex := index
	item := h.values[index]
	for kid := 2*index + 1; kid < len(h.values); kid = 2*index + 1 {
		right := kid + 1
		if right < len(h.values) {
			if h.values[kid] >= h.values[right] {
				kid = right
			}
		}
		h.values[index] = h.values[kid]
		index = kid
	}
	h.values[index] = item
	// Item was swapped into a leaf irrespective of whether that was right.
	// Now, move it back up where it should go.
	h.siftDown(originalIndex, index)
}

func run(input string) (int, error) {
	h := heap[int]{}
	for _, food := range strings.Split(input, "\n\n") {
		var calories int
		for _, line := range strings.Split(food, "\n") {
			c, err := strconv.Atoi(line)
			if err != nil {
				return 0, err
			}
			calories += c
		}
		h.push(-calories)
	}
	x, ok := h.pop()
	if !ok {
		return 0, fmt.Errorf("first element cannot be popped from the heap")
	}
	y, ok := h.pop()
	if !ok {
		return 0, fmt.Errorf("second element cannot be popped from the heap")
	}
	z, ok := h.pop()
	if !ok {
		return 0, fmt.Errorf("third element cannot be popped from the heap")
	}
	return -*x + -*y + -*z, nil
}

func main() {
	c, err := run(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
}
