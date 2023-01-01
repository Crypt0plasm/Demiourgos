package main

import (
  bloom "Demiourgos/Blooming"
  "fmt"
)

func main() {
  //CodingDivisionSetScanner()
  fmt.Println("")
  fmt.Println("====Snapshotting====")
  MultiChain, Len, Sum0 := bloom.CreateCodingDivisionChain()
  Owners := bloom.CreateCodingDivisionOwners(MultiChain)
  fmt.Println("A total of", Len, "have been snapshotted")
  fmt.Println("")
  fmt.Println("")

  //Raw Amount Computations
  fmt.Println("====Amount Computations====")
  All := bloom.CreateCodingDivisionAmountChain(Owners, MultiChain)
  SortedAll := bloom.SortBalanceSFTChain(All)
  AllSum := bloom.AddBalanceSFTChain(SortedAll)
  fmt.Println("Excluding SC, RAW (excluding SC, no single exceptions) SFTs sum is ", AllSum, "SFTs on", len(All), "Addresses")
  fmt.Println("")

  //Net Amount Computations (with single exceptions)
  AllException := bloom.CreateCodingDivisionAmountExceptionChain(SortedAll, true)
  SortedAllException := bloom.SortBalanceSFTChain(AllException)
  AllExceptionSum := bloom.AddBalanceSFTChain(SortedAllException)
  fmt.Println("Excluding SC, Net (excluding SC, with single exceptions) SFTs sum is ", AllExceptionSum, "SFTs on", len(AllException), "Addresses")
  fmt.Println("")

  //Raw Set Computations
  fmt.Println("====SET Computations====")
  Set := bloom.CreateCodingDivisionSetChain(Owners, MultiChain)
  SortedSet := bloom.SortBalanceSFTChain(Set)
  SetSum := bloom.AddBalanceSFTChain(SortedSet)
  fmt.Println("Excluding SC, RAW (excluding SC, no set exceptions) SFTs-SET sum is ", SetSum, "on ", len(SortedSet), " Addresses")
  fmt.Println("")

  //Net Set Computations
  SetException := bloom.CreateCodingDivisionSetExceptionChain(SortedSet, true)
  SortedSetException := bloom.SortBalanceSFTChain(SetException)
  SetExceptionSum := bloom.AddBalanceSFTChain(SortedSetException)
  fmt.Println("Excluding SC, Net (excluding SC, with set exceptions) SFTs-SET sum is ", SetExceptionSum, "on ", len(SortedSetException), " Addresses")
  fmt.Println("")

  //Reward Chain Computations (Amount Exception Set multiplied with reward 0.025)
  RewardChain := bloom.ComputeRewards(SortedAllException, "0.025")

  //Final Stats and file Outputs
  fmt.Println("====Total Stats====")
  fmt.Println("SFTs on Blockchain are ", Sum0, "on ", len(Owners), " Addresses")

  bloom.WriteListOneByOne("Output_All_Raw.txt", SortedAll)
  bloom.WriteListOneByOne("Output_All_Net.txt", SortedAllException)
  bloom.WriteListOneByOne("Output_Set_Raw.txt", SortedSet)
  bloom.WriteListOneByOne("Output_Set_Net.txt", SortedSetException)

  bloom.WriteListOneByOne("Output_All_Net_Reward.txt", RewardChain)
}
