package config

import (
	"fmt"
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Log struct {
		Level   string
		Fpath   string
		Msize   int
		Mage    int
		Mbackup int
	}

	Server struct {
		Port string
	}
}

func NewConfig(fpath string) *Config {
	c := new(Config)

	if file, err := os.Open(fpath); err != nil {
		panic(err)
		/* [코드리뷰]
		 * panic은 상당히 치명적인 에러를 마주한 경우에 발생합니다.
		 * panic을 manual하게 발생시키면 defer를 통해 예약해둔 함수는 호출을 보장할 순 있으나
		 * panic 이후 코드가 실행되지 못합니다. 프로그램이 종료됩니다.
		 * 어느 상황에서도 프로그램이 종료되는 경우는 막아야 합니다.
		 * 해당 코드는 message와 함께 err를 처리하는 코드를 추천드리고 싶습니다.
		 */
	} else {
		defer file.Close()
		//toml 파일 디코딩
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			fmt.Println(c)
			return c
		}
	}
	/* [코드리뷰]
	 * 해당 코드에는 하나의 function에서 간결한 이중 조건문이 발생하고 있습니다.
	 * 15 line으로 구성된 한눈에 들어오는 함수에서는 사용하기 적합합니다.
	 * 그러나 코드 라인수가 많아지고, 비즈니스 로직이 풍부해 지는 것을 고려한다면
	 * return으로 나갈 수 빠지는 case를 최소화하고, if문을 줄이는 방향으로 개발이 진행되어야 합니다. 
	 * 점진적으로 코드의 가독성이 조금 더 향상하게 됩니다.
	 * as-is: 
	 if file, err := os.Open(fpath); err != nil {
			panic(err)
		} else {
			defer file.Close()
			//toml 파일 디코딩
			if err := toml.NewDecoder(file).Decode(c); err != nil {
				panic(err)
			} else {
				fmt.Println(c)
				return c
			}
		}
	 * to-be:
	 if file, err := os.Open(fpath); err == nil {
			defer file.Close()
			//toml 파일 디코딩
			if err := toml.NewDecoder(file).Decode(c); err == nil {
				fmt.Println(c)
				return c, nil
			} 
		}
		return nil, err
	 */
}
