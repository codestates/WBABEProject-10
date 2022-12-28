package util

func PanicHandler(err error) {
	if err != nil {
		panic(err)
	}
}
/* [코드리뷰]
 * Handler라는 명칭을 통해 해당 function이 할 수 있는 기능을 정확하게 명시한 좋은 naming이라고 생각됩니다.
 * 그러나 Handler로서 수행되는 코드가 크게 복잡하지 않아, 단순 조건문으로 처리해도 괜찮은 코드입니다.
 * 일반적으로 현업에서 Handler를 사용하는 경우는, 특정 복잡한 상황을 처리해야 하는 경우를 하나의 파일을 handler로 구성하여,
 * 관련된 다양한 function들을 listing하는 방법을 사용해보시는 것을 추천드립니다.
 */
