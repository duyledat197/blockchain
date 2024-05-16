package blockchain_algorithm

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"os"

	"github.com/lmittmann/tint"
	"github.com/manifoldco/promptui"

	"openmyth/blockchain/pkg/blockchain"
	"openmyth/blockchain/pkg/blockchain/block"
	"openmyth/blockchain/pkg/blockchain/miner"
)

func getSelectTemplate(text string) *promptui.SelectTemplates {
	return &promptui.SelectTemplates{
		Active:   fmt.Sprintf("%s {{ . | underline | green }}", promptui.IconSelect),
		Label:    fmt.Sprintf("%s {{ . | blue }}: ", promptui.IconInitial),
		Selected: fmt.Sprintf("%s {{ . | white }}", promptui.IconGood+promptui.Styler(promptui.FGGreen)(fmt.Sprintf(" %s: ", text))),
	}
}

func getPromptTemplate() *promptui.PromptTemplates {
	return &promptui.PromptTemplates{
		Success: fmt.Sprintf("%s {{ . | green }}%s ", promptui.IconGood, promptui.Styler(promptui.FGGreen)(":")),
		Valid:   fmt.Sprintf("{{ . | blue }}%s ", promptui.Styler(promptui.FGBlue)(":")),
		Invalid: fmt.Sprintf("{{ . | blue }}%s ", promptui.Styler(promptui.FGBlue)(":")),
	}
}

const (
	// FeatureAddDataText represents for add data interaction
	FeatureAddDataText = "add data text"
	// FeatureListBlock represents for list block interaction
	FeatureListBlock = "list block"
	// FeatureExit represents for exit
	FeatureExit = "exit"
)

var (
	chooseFeature = promptui.Select{
		Label:     "Choose your interaction",
		Items:     []string{FeatureAddDataText, FeatureListBlock, FeatureExit},
		Templates: getSelectTemplate("action"),
	}

	addData = promptui.Prompt{
		Label:     "Please type text",
		Templates: getPromptTemplate(),
		Validate: func(str string) error {
			return nil
		},
	}
)

// Run runs the blockchain algorithm with the specified context.
//
// ctx: the context to run the blockchain algorithm
func Run(_ context.Context) {
	logger := slog.New(tint.NewHandler(os.Stdout, nil))

	slog.SetDefault(logger)

	var (
		numMiner  int
		broadcast = make(chan *block.Block)
	)
	flag.IntVar(&numMiner, "numMiner", 1, "number of miner")

	bc := blockchain.NewBlockchain()
	miners := make([]*miner.Miner, 0, numMiner)
	for range numMiner {
		miner := miner.NewMiner(*bc)
		miners = append(miners, miner)

		go miner.Mine(broadcast)
	}

	go func() {
		// broadcast to all miners
		for block := range broadcast {
			for _, miner := range miners {
				miner.AddBlockCh <- block
			}
		}
	}()

	for {
		_, feature, err := chooseFeature.Run()
		if err != nil {
			log.Fatalf("failed to select feature: %v", err)
		}
		isExit := false
		switch feature {
		case FeatureAddDataText:
			data, err := addData.Run()
			if err != nil {
				log.Printf("invalid data: %v", err)
			}

			for _, miner := range miners {
				// emit transactions to all nodes
				miner.TransactionCh <- []byte(data)
			}

		case FeatureListBlock:
			num := rand.Intn(numMiner)
			bc := miners[num].GetBlockChain()
			for _, block := range bc.GetBlocks() {
				log.Printf(`
					index: %d
					timestamp: %d
					data: %s
					previous block hash: %s
					hash: %s
					nonce: %d
				`, block.Index, block.Timestamp, block.Data, block.PrevBlockHash, block.Hash, block.Nonce)
			}
		case FeatureExit:
			isExit = true
		}
		if isExit {
			break
		}
	}
}
